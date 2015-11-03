/**
 * Created by Elvizlai on 2015/11/01 11:31
 * Copyright © PubCloud
 */

package controller
import (
	"github.com/ElvizLai/Blog/model/topic"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ElvizLai/Blog/enum"
	"strconv"
)

type TopicList struct {
	base
}

func (this *TopicList)TopicList() {
	this.TplNames = "topicList.html"

	currentPage := 1
	if page, err := strconv.Atoi(this.Input().Get("p")); err == nil {
		currentPage = page
	}

	topics, totalNum := topic.TopicList(currentPage)
	this.Data["topics"] = topics
	pagination.SetPaginator(this.Ctx, enum.CONST.PERPAGE, totalNum)
}