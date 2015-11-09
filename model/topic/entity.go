/**
 * Created by Elvizlai on 2015/10/31 22:43
 * Copyright © PubCloud
 */

package topic
import (
	"github.com/ElvizLai/Blog/model/user"
	"time"
)

type Topic struct {
	Id          int64
	User        *user.User `orm:"rel(fk)"`
	Title       string                    //topic title
	Tags        string                    //tags using ; split
	Abstract    string `orm:"type(text)"` //str before <!--more-->
	Markdown    string `orm:"type(text)"` //markdown
	HtmlContent string `orm:"type(text)"` //html
	PV          int64 `orm:"default(1)"`  //read count
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"null"`
	Previous    *Topic `orm:"-"`
	Next        *Topic `orm:"-"`
	Hash        string
}

//Ip与Path的联合唯一索引
func (u *Topic) TableUnique() [][]string {
	return [][]string{
		[]string{"Title", "Hash"},
	}
}