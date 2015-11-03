/**
 * Created by Elvizlai on 2015/11/01 12:38
 * Copyright © PubCloud
 */

package initial
import "github.com/astaxie/beego"

type errorController struct {
	beego.Controller
}

func setErrorPage()  {
	beego.ErrorController(&errorController{})
}

func (c *errorController) Error401() {
	c.TplNames = "base/401.html"
}

func (c *errorController) Error404() {
	c.TplNames = "base/404.html"
}

func (c *errorController) Error501() {
	c.Data["content"] = "server error"
	c.TplNames = "501.tpl"
}