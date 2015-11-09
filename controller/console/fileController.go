/**
 * Created by Elvizlai on 2015/11/01 22:02
 * Copyright © PubCloud
 */

package console
import (
	"io/ioutil"
	"github.com/ElvizLai/Blog/enum"
	"fmt"
	"os"
)

type FileController struct {
	base
}

func (this *FileController) Upload() {
	_, header, err := this.GetFile("editormd-image-file")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success":0, "message":err}
	}else {
		f, _ := header.Open()
		defer f.Close()
		data, _ := ioutil.ReadAll(f)

		filePath := enum.CONST.UPLOADPATH + "/" + fmt.Sprint(this.CurrentUser.Id) + "/"
		os.MkdirAll(filePath, os.ModePerm)

		ioutil.WriteFile(filePath + header.Filename, data, os.ModePerm)
		this.Data["json"] = map[string]interface{}{"success":1, "message":"OK", "url":"/" + filePath + header.Filename}
	}
	this.ServeJson()
}