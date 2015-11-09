/**
 * Created by elvizlai on 2015/11/9 16:07
 * Copyright © PubCloud
 */
package controller
import (
	"github.com/ElvizLai/Blog/model/topic"
	"strconv"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ElvizLai/Blog/enum"
)

type TagController struct {
	base
}

func (this *TagController) TopicList() {
	this.TplNames = "tag.html"

	tag := this.Ctx.Input.Param(":tag")

	currentPage := 1
	if page, err := strconv.Atoi(this.Input().Get("p")); err == nil {
		currentPage = page
	}

	topics, totalNum := topic.GetTopicListByTag(currentPage, tag)
	this.Data["topics"] = topics
	pagination.SetPaginator(this.Ctx, enum.CONST.PERPAGE, totalNum)
}
