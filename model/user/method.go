/**
 * Created by Elvizlai on 2015/10/31 22:41
 * Copyright © PubCloud
 */

package user
import (
	"github.com/ElvizLai/Blog/util"
	"github.com/astaxie/beego/orm"
)

//create an user, with dup_key error for email or nickname.
func AddUser(email, nickname, password string) error {
	salt := util.RandString(8)
	password = util.Md5(salt + password)
	u := &User{Email:email, NickName:nickname, Salt:salt, Password:password}

	o := orm.NewOrm()
	_, err := o.Insert(u)

	return err
}

func GetUserByEmail(email string) *User {
	user := &User{Email:email}
	if orm.NewOrm().Read(user, "Email") != nil {
		user = nil
	}
	return user
}