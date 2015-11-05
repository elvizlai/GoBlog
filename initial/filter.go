/**
 * Created by Elvizlai on 2015/11/01 13:42
 * Copyright © PubCloud
 */

package initial
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//拦截没有鉴权的请求
func setFilter() {
	beego.InsertFilter("/console/*", beego.BeforeRouter, userFilter)
}

var userFilter = func(ctx *context.Context) {
	if ctx.Input.Session("user") == nil {
		if ctx.Input.Method() == "GET" {
			ctx.Redirect(302, "/login")
		}else {
			ctx.Abort(401, "401")
		}
	}
}