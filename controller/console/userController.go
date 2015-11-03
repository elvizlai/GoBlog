/**
 * Created by Elvizlai on 2015/10/31 22:51
 * Copyright © PubCloud
 */

package console
import (
	"github.com/ElvizLai/Blog/model/user"
	"github.com/astaxie/beego/validation"
	"github.com/ElvizLai/Blog/enum"
	"github.com/ElvizLai/Blog/util"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	base
}

func (this *UserController) Login() {
	if this.Ctx.Input.Method() == "GET" {
		if this.GetSession("user") != nil {
			this.Redirect("/", 302)
		}
		this.TplNames = "console/login.html"
		return
	}

	json := this.ReqJson()
	email := json.Get("email").MustString()
	password := json.Get("password").MustString()

	valid := validation.Validation{}
	valid.Email(email, "email")
	valid.MinSize(password, 6, "passwordMin")
	valid.MaxSize(password, 12, "passwordMax")

	if valid.HasErrors() {
		this.CustomAbort(enum.RespCode.BadRequest.Code(), enum.RespCode.BadRequest.Str())
	}

	user := user.GetUserByEmail(email)
	if user == nil {
		//用户不存在
		this.RespJson(enum.RespCode.UserNotExist, nil)
	}else if util.Md5(user.Salt + password) != user.Password {
		//密码错误
		this.RespJson(enum.RespCode.PasswordIncorrect, nil)
	}else {
		this.SetSession("user", user)
		this.RespJson(enum.RespCode.OK, map[string]interface{}{"url":"/"})
	}
}

func (this *UserController) Register() {
	if this.Ctx.Input.Method() == "GET" {
		if can, err := beego.AppConfig.Bool("CanBeRegister"); err == nil && can {
			this.Data["CanBeRegister"] = true
		}
		this.TplNames = "console/register.html"
		return
	}

	req := this.ReqJson()
	email := req.Get("email").MustString()
	nickName := req.Get("nickName").MustString()
	password := req.Get("password").MustString()

	valid := validation.Validation{}
	valid.Email(email, "email")
	valid.MinSize(nickName, 6, "nickNameMin")
	valid.MaxSize(nickName, 12, "nickNameMax")
	valid.MinSize(password, 6, "passwordMin")
	valid.MaxSize(password, 12, "passwordMax")

	if valid.HasErrors() {
		this.CustomAbort(enum.RespCode.BadRequest.Code(), enum.RespCode.BadRequest.Str())
	}

	err := user.AddUser(email, nickName, password)
	if err == nil {
		this.RespJson(enum.RespCode.OK, nil)
	}else {
		if strings.Contains(err.Error(), "email") {
			this.RespJson(enum.RespCode.EmailExist, nil)
		}else if strings.Contains(err.Error(), "nick_name") {
			this.RespJson(enum.RespCode.NickNameExist, nil)
		}else {
			beego.Error(err)
		}
	}
}