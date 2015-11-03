/**
 * Created by Elvizlai on 2015/11/03 19:54
 * Copyright © PubCloud
 */

package visitLog
import "time"

type VisitLog struct {
	Id         int64
	Ip         string
	Path       string
	Lat        float64 `orm:"digits(8);decimals(4)"`
	Lng        float64 `orm:"digits(8);decimals(4)"`
	City       string
	Feq        int64 `orm:"default(1)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

//Ip与Path的联合唯一索引
func (u *VisitLog) TableUnique() [][]string {
	return [][]string{
		[]string{"Ip", "Path"},
	}
}
