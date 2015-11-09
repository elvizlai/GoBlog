/**
 * Created by elvizlai on 2015/11/9 16:18
 * Copyright © PubCloud
 */
package controller
import (
	"github.com/ElvizLai/Blog/model/topic"
	"strconv"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ElvizLai/Blog/enum"
)

type TopicController struct {
	base
}

func (this *TopicController)TopicList() {
	this.TplNames = "topicList.html"

	currentPage := 1
	if page, err := strconv.Atoi(this.Input().Get("p")); err == nil {
		currentPage = page
	}

	topics, totalNum := topic.GetTopicList(currentPage)
	this.Data["topics"] = topics
	pagination.SetPaginator(this.Ctx, enum.CONST.PERPAGE, totalNum)
}

func (this *TopicController) TopicDetail() {
	this.TplNames = "topicDetail.html"

	id:= this.Ctx.Input.Param(":id")
	tp := topic.GetTopicById(id)

	if tp == nil {
		this.Abort("404")
	}else {
		if this.CurrentUser != nil && this.CurrentUser.Id == tp.User.Id {
			this.Data["canModify"] = true
		}

		topic.AddPV(id)//pv+1

		this.Data["topic"] = tp
	}
}