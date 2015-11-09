/**
 * Created by Elvizlai on 2015/11/01 10:41
 * Copyright © PubCloud
 */

package console
import (
	"github.com/ElvizLai/Blog/enum"
	"github.com/ElvizLai/Blog/model/topic"
	"strings"
	"regexp"
	"fmt"
)

type TopicController struct {
	base
}

func (this *TopicController) NewTopic() {
	if this.Ctx.Input.Method() == "GET" {
		this.TplNames = "console/new_topic.html"
		return
	}

	req := this.ReqJson()

	title := req.Get("title").MustString()
	tags := req.Get("tags").MustString()

	if tags == "" {
		tags = "未分类"
	}else {
		tags = strings.TrimSpace(tags)
		tags = strings.Replace(tags, "；", ";", -1)
	}

	markdown := req.Get("markdown").MustString()
	htmlContent := req.Get("htmlContent").MustString()
	abstract := getAbstract(htmlContent)

	err := topic.AddTopic(this.CurrentUser, title, tags, abstract, markdown, htmlContent)

	if err == nil {
		this.RespJson(enum.RespCode.OK, nil)
	}else {
		this.RespJson(enum.RespCode.Conflict, err)
	}
}

func (this *TopicController) ModifyTopic() {
	idStr := this.Ctx.Input.Param(":id")

	if this.Ctx.Input.Method() == "GET" {
		this.TplNames = "console/modify_topic.html"
		t := topic.GetTopicById(idStr)
		if t == nil {
			this.Abort("404")
		}else {
			fmt.Println(this.CurrentUser)
			if t.User.Id != this.CurrentUser.Id {
				this.Redirect("/topic/" + idStr, 302)
				return
			}
			this.Data["topic"] = t
		}
		return
	}

	req := this.ReqJson()
	method := req.Get("method").MustString()
	var err error
	if method == "update" {
		title := req.Get("title").MustString()
		tags := req.Get("tags").MustString()
		hash := req.Get("hash").MustString()

		if tags == "" {
			tags = "未分类"
		}

		markdown := req.Get("markdown").MustString()
		htmlContent := req.Get("htmlContent").MustString()
		abstract := getAbstract(htmlContent)

		err = topic.ModifyTopic(idStr, this.CurrentUser, title, tags, abstract, markdown, htmlContent, hash)
	}else if method == "delete" {
		err = topic.DeleteTopic(idStr, this.CurrentUser)
	}else {
		this.RespJson(enum.RespCode.BadRequest, nil)
	}

	if err == nil {
		this.RespJson(enum.RespCode.OK, nil)
	}else {
		this.RespJson(enum.RespCode.Conflict, err)
	}

}

func getAbstract(markdown string) string {
	reg := regexp.MustCompile(`<!-{2,}more-{2,}>`)
	index := reg.FindStringIndex(markdown)
	abstract := ""
	if index != nil {
		abstract = markdown[:index[0]]
	}else {
		abstract = markdown
	}
	return abstract
}