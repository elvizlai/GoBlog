/**
 * Created by Elvizlai on 2015/10/31 22:26
 * Copyright © PubCloud
 */

package initial
import (
	"github.com/astaxie/beego"
)

func setLogger()  {
	beego.SetLogger("file", `{"filename":"log"}`)
	//beego.SetLevel(beego.LevelInformational)
}