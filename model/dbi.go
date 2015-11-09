/**
 * Created by Elvizlai on 2015/10/31 22:30
 * Copyright © PubCloud
 */

package model
import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/ElvizLai/Blog/model/user"
	"github.com/ElvizLai/Blog/enum"
	"github.com/ElvizLai/Blog/model/topic"
	"github.com/ElvizLai/Blog/model/visitLog"
)

func init() {
	if beego.RunMode != "prod" {
		orm.Debug = true
	}

	if err := orm.RegisterDataBase("default", "sqlite3", enum.CONST.DBNAME, 50); err != nil {
		panic(err)
	}

	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(topic.Topic))
	orm.RegisterModel(new(visitLog.VisitLog))

	orm.RunSyncdb("default", false, false)
}