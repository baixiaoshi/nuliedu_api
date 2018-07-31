package qiniu

import (
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"fmt"
	"context"
)

type Upload struct {
	AccessKey string `json:"accesskey"`
	SecretKey string `json:"secretkey"`
}

func NewUpload(accessKey, secreKey string) *Upload {
	return &Upload{AccessKey: accessKey, SecretKey: secreKey}
}


// @param bucket 七牛后台bucket名称
// @param key 七牛后台秘钥
// @param filepath 文件本地绝对路径
// @param filepath 文件本地绝对路径
func (u *Upload) Upload(bucket, key, filepath string) (*storage.PutRet, error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(u.AccessKey, u.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filepath, &putExtra)
	if err != nil {
		return nil, err
	}
	fmt.Println(ret.Key, ret.Hash)

	return &ret, nil
}
