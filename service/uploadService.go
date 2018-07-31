package service

import (
	"nuliedu_api/models"
	"nuliedu_api/lib/qiniu"
	"github.com/astaxie/beego"
)

type UploadService struct {

}


type uploadInfo struct {
	//数据库中的自增id
	Id int
	//本地保存地址
	ImgPath string
	//七牛path
	QiNiuPath string

}

func (u *UploadService) Upload(savePath string, size int) (*uploadInfo, error) {

	//上传到七牛云
	accessKey := beego.AppConfig.String("qiniu.access_key")
	secreKey := beego.AppConfig.String("qiniu.secre_key")
	qiniuUpload := qiniu.NewUpload(accessKey, secreKey)
	qiniuUpload.Upload("", "", savePath)

	//插入到数据库

	m := models.Imgs{}
	m.ImgPath  = savePath
	m.QiNiuPath = ""
	m.ImgWidth =0
	m.ImgHeight = 0
	m.ImgSize = size



}


