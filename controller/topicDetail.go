/**
 * Created by Elvizlai on 2015/11/01 12:00
 * Copyright © PubCloud
 */

package controller
import (
	"github.com/ElvizLai/Blog/model/topic"
	"strconv"
)

type TopicDetail struct {
	base
}

func (this *TopicDetail) TopicDetail() {
	this.TplNames = "topicDetail.html"

	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
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