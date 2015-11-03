/**
 * Created by Elvizlai on 2015/11/01 22:44
 * Copyright © PubCloud
 */

package controller
import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)

type FileHandler struct {
	base
}

func (this *FileHandler) Get() {
	filePath := strings.TrimLeft(this.Ctx.Request.RequestURI, "/")
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		//todo 找不到图片
	}

	data, _ := ioutil.ReadFile(filePath)

	this.Ctx.ResponseWriter.Header().Add("Cache-Control", "public, max-age=31536000")
	//在资源存在的情况下，为所有的请求返回Last-Modified头
	this.Ctx.ResponseWriter.Header().Add("Last-Modified", fmt.Sprint(fileInfo.ModTime().Unix()))
	this.Ctx.ResponseWriter.Write(data)
}