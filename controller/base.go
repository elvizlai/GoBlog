/**
 * Created by Elvizlai on 2015/10/31 22:52
 * Copyright © PubCloud
 */

package controller
import (
	"github.com/astaxie/beego"
	"github.com/ElvizLai/Blog/model/user"
	"strings"
	"github.com/ElvizLai/Blog/model/visitLog"
	"github.com/astaxie/beego/context"
)

type base struct {
	beego.Controller
	CurrentUser *user.User
}

func (this *base) Prepare() {
	this.Layout = "Layout.html"
	//是否已登录
	if s := this.GetSession("user"); s != nil {
		this.CurrentUser = s.(*user.User)
		this.Data["isLogin"] = true
	}

	//访问统计 过滤掉files目录
	go func(input *context.BeegoInput) {
		path := input.Url()
		if !strings.HasPrefix(path, "/files/") {
			visitLog.AddVisitLog(input.IP(), path)
		}
	}(this.Ctx.Input)

}