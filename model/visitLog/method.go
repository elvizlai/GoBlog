/**
 * Created by Elvizlai on 2015/11/03 19:54
 * Copyright © PubCloud
 */

package visitLog
import (
	"github.com/astaxie/beego/orm"
	"github.com/ElvizLai/Blog/util"
	"github.com/astaxie/beego"
)

//添加访问记录
func AddVisitLog(ip, path string) {
	vl := &VisitLog{Ip:ip, Path:path, Feq:1}

	o := orm.NewOrm()

	//已存在
	if o.Read(vl, "Ip", "Path") == nil {
		//更新feq+1
		vl.Feq += 1
		if _, err := o.Update(vl); err != nil {
			beego.Error(err)
		}
		return
	}

	//不存在
	if j := util.InfoGeoByIP(ip); j != nil {
		vl.City = j.Get("regionName").MustString() + "," + j.Get("city").MustString()
		vl.Lng = j.Get("lon").MustFloat64()
		vl.Lat = j.Get("lat").MustFloat64()
	}

	if _, err := o.Insert(vl); err != nil {
		beego.Error(err)
	}
}