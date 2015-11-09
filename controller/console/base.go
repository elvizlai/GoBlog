/**
 * Created by Elvizlai on 2015/10/31 22:52
 * Copyright © PubCloud
 */

package console
import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"github.com/ElvizLai/Blog/enum"
	"github.com/ElvizLai/Blog/model/user"
)

type base struct {
	beego.Controller
	CurrentUser *user.User
}

func (this *base) Prepare() {
	this.Layout = "console/Layout.html"
	//get current user from session
	if s := this.GetSession("user"); s != nil {
		this.CurrentUser = s.(*user.User)
	}
}

func (this *base) ReqJson() *simplejson.Json {
	defer func() {
		if err := recover(); err != nil && err.(string) == "jsonErr" {
			this.CustomAbort(enum.RespCode.BadRequest.Code(), enum.RespCode.BadRequest.Str())
		}
	}()

	if json, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		return json
	}else {
		panic("jsonErr")
	}
}

func (this *base) RespJson(e *enum.Code, result interface{}) {
	this.Data["json"] = map[string]interface{}{"code":e.Code(), "msg":e.Str(), "result":result}
	this.ServeJson()
}