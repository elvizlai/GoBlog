/**
 * Created by Elvizlai on 2015/11/01 20:44
 * Copyright © PubCloud
 */

package controller

type About struct {
	base
}

func (this *About) Get() {
	this.TplNames = "about.html"
}