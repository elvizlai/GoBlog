/**
 * Created by Elvizlai on 2015/10/31 22:26
 * Copyright © PubCloud
 */

package initial
import (
	"github.com/astaxie/beego"
	"os"
)

func setLogger() {
	//check if logs dir exist, if not create it
	_, err := os.Stat("logs")
	if err != nil {
		os.Mkdir("logs", os.ModePerm)
	}
	beego.SetLogger("file", `{"filename":"logs/log"}`)
	if beego.RunMode == "prod" {
		//setting log level to info, ignore debug
		beego.SetLevel(beego.LevelInformational)
	}
}