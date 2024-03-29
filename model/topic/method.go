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
	"time"
	"errors"
)

func AddTopic(user *user.User, title, tags, abstract, markdown, htmlContent string) error {
	topic := &Topic{User:user, Title:title, Tags:tags, Abstract:abstract, Markdown:markdown, HtmlContent:htmlContent, PV:1}
	topic.Hash = util.Md5(fmt.Sprint(topic))
	_, err := orm.NewOrm().Insert(topic)
	return err
}

func GetTopicList(page int) ([]Topic, int64) {
	o := orm.NewOrm()
	topics := []Topic{}

	totalNum, _ := o.QueryTable("Topic").Count()
	o.QueryTable("Topic").OrderBy("-CreateTime").Offset(enum.CONST.PERPAGE * (page - 1)).Limit(enum.CONST.PERPAGE).All(&topics)

	for i, l := 0, len(topics); i < l; i++ {
		topics[i].CreateTime = topics[i].CreateTime.In(enum.CONST.TIMEZONE)
	}

	return topics, totalNum
}

func GetTopicListByTag(page int, tag string) ([]Topic, int64) {
	o := orm.NewOrm()
	topics := []Topic{}
	totalNum, _ := o.QueryTable("Topic").Filter("tags__icontains", tag).Count()
	o.QueryTable("Topic").Filter("tags__icontains", tag).OrderBy("-CreateTime").Offset(enum.CONST.PERPAGE * (page - 1)).Limit(enum.CONST.PERPAGE).All(&topics)

	for i, l := 0, len(topics); i < l; i++ {
		topics[i].CreateTime = topics[i].CreateTime.In(enum.CONST.TIMEZONE)
	}

	return topics, totalNum
}

func GetTopicById(id interface{}) *Topic {
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

func AddPV(id interface{}) {
	o := orm.NewOrm()
	o.QueryTable("Topic").Filter("Id", id).Update(orm.Params{"PV":orm.ColValue(orm.Col_Add, 1)})
}

func ModifyTopic(id interface{}, u *user.User, title, tags, abstract, markdown, htmlContent, hash string) error {
	//1、判断是否真的为该用户 2、hash值是否匹配
	o := orm.NewOrm()
	t := &Topic{}
	o.QueryTable("Topic").Filter("Id", id).RelatedSel().One(t)
	if t.User.Id != u.Id {
		return errors.New(enum.RespCode.UnAuthorized.Str())
	}

	if t.Hash != hash {
		return errors.New(enum.RespCode.Conflict.Str())
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

func DeleteTopic(id interface{}, u *user.User) error {
	o := orm.NewOrm()
	t := &Topic{}
	o.QueryTable("Topic").Filter("Id", id).RelatedSel().One(t)
	if t.User.Id != u.Id {
		return errors.New(enum.RespCode.UnAuthorized.Str())
	}

	_, err := o.Delete(t)

	//todo backup
	return err
}