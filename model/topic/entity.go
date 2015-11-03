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
	Title       string                    //标题
	Tags        string                    //标签，使用符号;分隔
	Abstract    string `orm:"type(text)"` //摘要
	Markdown    string `orm:"type(text)"` //markdown正文
	HtmlContent string `orm:"type(text)"` //html正文
	PV          int64 `orm:"default(1)"`  //浏览量
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"null"`
	Previous    *Topic `orm:"-"`
	Next        *Topic `orm:"-"`
	Hash        string                    //用于校验非法修改
}
