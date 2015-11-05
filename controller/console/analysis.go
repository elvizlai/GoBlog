/**
 * Created by Elvizlai on 2015/11/05 19:39
 * Copyright © PubCloud
 */

package console
import (
	"github.com/astaxie/beego/orm"
	"github.com/ElvizLai/Blog/model/visitLog"
)

type MapController struct {
	base
}

func (this *MapController) Get() {
	this.Layout = ""
	this.TplNames = "console/map.html"
}

func (this *MapController) Post() {
	o := orm.NewOrm()
	sql := `SELECT city, lat, lng, SUM(feq) as feq FROM visit_log GROUP BY lat,lng`
	visiters := []visitLog.VisitLog{}
	o.Raw(sql).QueryRows(&visiters)
	geoCoord := map[string][2]float64{}
	data := []map[string]interface{}{}
	for i, l := 0, len(visiters); i < l; i++ {
		geoCoord[visiters[i].City] = [2]float64{visiters[i].Lng, visiters[i].Lat}
		data = append(data, map[string]interface{}{"name":visiters[i].City, "value":visiters[i].Feq})
	}

	this.Data["json"] = map[string]interface{}{"data":data, "geoCoord":geoCoord}
	this.ServeJson()
}
