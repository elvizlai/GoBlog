/**
 * Created by Elvizlai on 2015/10/31 22:43
 * Copyright © PubCloud
 */

package topic
import (
	"github.com/ElvizLai/Blog/model/user"
	"github.com/astaxie/beego/orm"
	"github.com/ElvizLai/Blog/enum"
	"github.com/ElvizLai/Blog/util"
	"fmt"
	"github.com/juju/errors"
	"time"
)

func AddTopic(user *user.User, title, tags, abstract, markdown, htmlContent string) error {
	topic := &Topic{User:user, Title:title, Tags:tags, Abstract:abstract, Markdown:markdown, HtmlContent:htmlContent, PV:1}
	topic.Hash = util.Md5(fmt.Sprint(topic))
	_, err := orm.NewOrm().Insert(topic)
	return err
}

func TopicList(page int) ([]Topic, int64) {
	o := orm.NewOrm()
	topics := []Topic{}

	totalNum, _ := o.QueryTable("Topic").Count()
	o.QueryTable("Topic").OrderBy("-CreateTime").Offset(enum.CONST.PERPAGE * (page - 1)).Limit(enum.CONST.PERPAGE).All(&topics)

	for i, l := 0, len(topics); i < l; i++ {
		topics[i].CreateTime = topics[i].CreateTime.In(enum.CONST.TIMEZONE)
	}

	return topics, totalNum
}

func GetTopicById(id int64) *Topic {
	o := orm.NewOrm()
	topic := &Topic{}
	if o.QueryTable("Topic").Filter("Id", id).RelatedSel().One(topic) != nil {
		//不存在该文章
		return nil
	}

	//上一篇与下一篇查询
	topics := []Topic{}
	count, _ := o.QueryTable("Topic").Filter("Id__gt", topic.Id).OrderBy("Id").Limit(1).All(&topics)
	if count == 1 {
		topic.Previous = &topics[0]
	}

	count, _ = o.QueryTable("Topic").Filter("Id__lt", topic.Id).OrderBy("-Id").Limit(1).All(&topics)
	if count == 1 {
		topic.Next = &topics[0]
	}

	topic.CreateTime = topic.CreateTime.In(enum.CONST.TIMEZONE)
	return topic
}

func AddPV(id int64) {
	o := orm.NewOrm()
	o.QueryTable("Topic").Filter("Id", id).Update(orm.Params{"PV":orm.ColValue(orm.Col_Add, 1)})
}

func ModifyTopic(id int64, u *user.User, title, tags, abstract, markdown, htmlContent, hash string) error {
	//1、判断是否真的为该用户 2、hash值是否匹配
	o := orm.NewOrm()
	t := &Topic{}
	o.QueryTable("Topic").Filter("Id", id).RelatedSel().One(t)
	if t.User.Id != u.Id {
		return errors.New("aaaaa")
	}

	if t.Hash != hash {
		return errors.New("bbbbb")
	}

	t.Title = title
	t.Tags = tags
	t.Abstract = abstract
	t.Markdown = markdown
	t.HtmlContent = htmlContent
	t.Hash = util.Md5(fmt.Sprint(t))
	t.UpdateTime = time.Now()
	_, err := o.Update(t)
	return err
}

func DeleteTopic(id int64, u *user.User) error {
	o := orm.NewOrm()
	t := &Topic{}
	o.QueryTable("Topic").Filter("Id", id).RelatedSel().One(t)
	if t.User.Id != u.Id {
		return errors.New("aaaaa")
	}

	_, err := o.Delete(t)
	return err
}