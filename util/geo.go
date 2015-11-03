/**
 * Created by Elvizlai on 2015/10/31 22:26
 * Copyright © PubCloud
 */

package util
import (
	"github.com/astaxie/beego/httplib"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/astaxie/beego"
)

const reqUrl = "http://ip-api.com/json/%s"

func InfoGeoByIP(ip string) *simplejson.Json {
	req := httplib.Get(fmt.Sprintf(reqUrl, ip))
	resp, err := req.Bytes()
	if err != nil {
		beego.Error(err)
		return nil
	}

	j, err := simplejson.NewJson(resp)
	if err != nil {
		beego.Error(err)
		return nil
	}

	status := j.Get("status").MustString()
	if status == "success" {
		return j
	}else {
		beego.Error(err)
		return nil
	}
}