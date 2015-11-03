/**
 * Created by Elvizlai on 2015/11/01 20:44
 * Copyright © PubCloud
 */

package controller

type Archives struct {
	base
}

func (this *Archives) Get()  {
	this.TplNames = "archives.html"
}