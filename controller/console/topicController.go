/**
 * Created by Elvizlai on 2015/11/01 10:41
 * Copyright © PubCloud
 */

package console
import (
	"github.com/ElvizLai/Blog/enum"
	"github.com/ElvizLai/Blog/model/topic"
	"github.com/ElvizLai/Blog/model/user"
	"strings"
	"regexp"
	"strconv"
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

	u, _ := this.GetSession("user").(*user.User)

	err := topic.AddTopic(u, title, tags, abstract, markdown, htmlContent)//.AddArticle(this.CurrentUser.Id, title, tags, markdown, htmlContent)

	if err == nil {
		this.RespJson(enum.RespCode.OK, nil)
	}else {
		//this.RespJson(enum.UNKNOWN, err)
	}
}

func (this *TopicController) ModifyTopic() {
	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	u, _ := this.GetSession("user").(*user.User)
	if this.Ctx.Input.Method() == "GET" {
		this.TplNames = "console/modify_topic.html"
		t := topic.GetTopicById(id)
		if t == nil {
			this.Abort("404")
		}else {
			if t.User.Id != u.Id {
				this.Redirect("/topic/" + fmt.Sprint(id), 302)
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

		err = topic.ModifyTopic(id, u, title, tags, abstract, markdown, htmlContent, hash)
	}else if method == "delete" {
		err = topic.DeleteTopic(id,u)
	}else {
		this.RespJson(enum.RespCode.BadRequest, nil)
	}

	if err == nil {
		this.RespJson(enum.RespCode.OK, nil)
	}else {
		this.RespJson(enum.RespCode.BadRequest, err)
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