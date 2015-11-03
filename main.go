/**
 * Created by Elvizlai on 2015/10/31 22:26
 * Copyright © PubCloud
 */

package main
import (
	_ "github.com/ElvizLai/Blog/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Info(beego.VERSION)

	beego.Run()
}
