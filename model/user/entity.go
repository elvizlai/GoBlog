/**
 * Created by Elvizlai on 2015/10/31 22:41
 * Copyright © PubCloud
 */

package user
import (
	"time"
)

type User struct {
	Id         int64
	Email      string `orm:"unique"`
	NickName   string `orm:"unique"`
	Salt       string
	Password   string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` //user create time
}
