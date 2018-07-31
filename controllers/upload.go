package controllers

import (
	"fmt"
	"image"
	"nuliedu_api/util"
	util2 "github.com/augneb/util"
	"nuliedu_api/helper"
	"github.com/astaxie/beego"
	"time"
	"path/filepath"
	"nuliedu_api/service"
)

type UploadController struct {
	BaseController
}


//自己上传图片，这里不做任何校验
//上传后获取需要的信息
// 保存到相应的服务器位置
// 上传到七牛云上面
// 返回上传的图片信息,id,qiniu_path
func (u *UploadController) Img() {

	_, h, err := u.GetFile("file") //获取上传的文件
	if err != nil {
		u.Response(false, "upload err:" + err.Error(), nil)
		return
	}

	fileExt := filepath.Ext(h.Filename)
	//创建目录并保存图片
	saveDir := beego.AppConfig.String("upload.image_save_path")
	dateStr := time.Now().Format("2006-01-02")
	t := time.Now().UnixNano()
	imgsize := h.Size

	abImgPath := fmt.Sprintf("%s/%s/%d.%s", saveDir, dateStr, t, fileExt)

	err = helper.CreateDir(abImgPath)
	if err != nil {
		u.Response(false, "create dir fail:"  + err.Error(), nil)
		return
	}

	//保存图片
	err = u.SaveToFile("file", abImgPath)
	if err != nil {
		u.Response(false, "save fail:" + err.Error(), nil)
		return
	}

	//上传
	uploadService := service.UploadService{}
	upInfo, err := uploadService.Upload(abImgPath, int(imgsize))
	if err != nil {
		u.Response(false, "upload err:" + err.Error(), nil)
		return
	}


	u.Response(true, "upload ok", upInfo)
}
