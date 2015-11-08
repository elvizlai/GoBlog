/**
 * Created by Elvizlai on 2015/11/07 22:44
 * Copyright © PubCloud
 */

package backup
import (
	"github.com/astaxie/beego/httplib"
	"fmt"
	"io/ioutil"
	"github.com/astaxie/beego"
	"time"
	"os"
	"io"
	"archive/zip"
	"github.com/ElvizLai/Blog/enum"
	"path/filepath"
	"net/http"
)

var url = "https://content.dropboxapi.com/2/files/upload"
var argF = `{"path": "/Backup/%s/%s","mode": "add","autorename": false,"mute": false}`
var dropboxAuthorization = ""


func init() {
	//check if DropboxAuthorization exist
	if dropboxAuthorization = beego.AppConfig.String("DropboxAuthorization"); dropboxAuthorization == "" {
		return
	}

	today := time.Now().Day()
	go func() {
		for {
			if now := time.Now(); now.Day() != today {
				beego.Info("BackUp begin at:", now)
				backUp()
				today = now.Day()
			}
			<-time.After(time.Hour)
		}
	}()
}

//comp file
func backUp() {
	backUpFile := beego.AppName + "_" + time.Now().Format("060102") + ".zip"
	fw, err := os.Create(backUpFile)
	if err != nil {
		beego.Error(err)
		return
	}
	defer fw.Close()

	zw := zip.NewWriter(fw)
	defer zw.Close()

	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		src, _ := os.Open(path)
		defer src.Close()
		h := &zip.FileHeader{Name: path, Method: zip.Deflate, Flags: 0x800}
		fileName, _ := zw.CreateHeader(h)
		io.Copy(fileName, src)
		zw.Flush()
		return nil
	}

	if err := filepath.Walk(enum.CONST.DBNAME, walk); err != nil {
		beego.Error(err)
		return
	}

	if err := filepath.Walk(enum.CONST.FILEPATH, walk); err != nil {
		beego.Error(err)
		return
	}

	//file upload
	req := httplib.Post(url)
	req.Header("Authorization", dropboxAuthorization)
	req.Header("Dropbox-API-Arg", fmt.Sprintf(argF, beego.AppName, backUpFile))
	req.Header("Content-Type", "application/octet-stream")

	data, err := ioutil.ReadFile(backUpFile)
	if err != nil {
		beego.Error(err)
		return
	}
	req.Body(data)

	resp, err := req.SendOut()
	if err != nil || resp.StatusCode != http.StatusOK {
		beego.Error(err, resp)
	}

	//file delete
	os.Remove(backUpFile)
}