/**
 * Created by Elvizlai on 2015/10/31 23:22
 * Copyright © PubCloud
 */

package enum
import (
	"time"
	"github.com/astaxie/beego"
)

var CONST *c

type c struct {
	DBNAME     string
	PERPAGE    int
	TIMEZONE   *time.Location
	UPLOADPATH string
}

func init() {
	var err error
	CONST = new(c)
	CONST.DBNAME = "data.sqlite"
	CONST.UPLOADPATH = "files"
	CONST.PERPAGE = 10//todo read from config

	CONST.TIMEZONE, err = time.LoadLocation(beego.AppConfig.String("TimeZone"))
	if err != nil {
		CONST.TIMEZONE, _ = time.LoadLocation("Asia/Shanghai")
	}
}