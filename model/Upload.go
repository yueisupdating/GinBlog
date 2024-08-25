package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"log"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}

func setConfig() storage.Config {
	cfg := storage.Config{
		Zone:          selectZone(utils.Zone),
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	return cfg
}

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	mac := auth.New(utils.AccessKey, utils.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope:   utils.Bucket,
		Expires: 7200,
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := setConfig()
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		log.Panicln(err)
		return "", errmsg.ERROR
	}
	return utils.QiniuSever + ret.Key, errmsg.SUCCESS
}
