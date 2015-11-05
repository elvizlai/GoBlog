/**
 * Created by Elvizlai on 2015/10/31 22:29
 * Copyright © PubCloud
 */

package router
import (
	"github.com/astaxie/beego"
	"github.com/ElvizLai/Blog/controller/console"
	"github.com/ElvizLai/Blog/controller"
	"github.com/ElvizLai/Blog/enum"
)

func init() {
	beego.Router("/", &controller.TopicList{}, "get:TopicList")
	beego.Router("/topic/:id([0-9]+)", &controller.TopicDetail{}, "get:TopicDetail")
	beego.Router("/archives", &controller.Archives{})
	beego.Router("/about", &controller.About{})
	beego.Router("/" + enum.CONST.FILEPATH + "/:id([0-9]+)/*", &controller.FileHandler{})

	beego.Router("/login", &console.UserController{}, "get,post:Login")
	beego.Router("/register", &console.UserController{}, "get,post:Register")
	beego.Router("/console/newTopic", &console.TopicController{}, "get,post:NewTopic")
	beego.Router("/console/modifyTopic/:id([0-9]+)", &console.TopicController{}, "get,post:ModifyTopic")
	beego.Router("/console/map", &console.MapController{})
	beego.Router("/upload", &console.FileController{}, "post:Upload")
}