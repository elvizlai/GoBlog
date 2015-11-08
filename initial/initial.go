/**
 * Created by Elvizlai on 2015/10/31 22:31
 * Copyright © PubCloud
 */

package initial
import (
	"os"
	"github.com/ElvizLai/Blog/enum"
	_ "github.com/ElvizLai/Blog/router"
	_ "github.com/ElvizLai/Blog/model"
	_ "github.com/ElvizLai/Blog/backup"
	"github.com/astaxie/beego"
)

func init() {
	beego.SessionOn = true
	beego.SessionName = "token"
	beego.CopyRequestBody = true

	//setting Logger
	setLogger()
	//setting error page
	setErrorPage()
	//setting filter
	setFilter()

	//remove sqlite journal
	os.Remove(enum.CONST.DBNAME + "-journal")

	//create file
	os.Mkdir(enum.CONST.FILEPATH, os.ModePerm)

	//adding function fot html template
	registerFunc()
}