package oss

import (
	"bytes"
	"context"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadFile(fileData []byte, fileSize int64) (string, error) {
	server := conf.GetConf().OSS

	// 将 []byte 转换为 io.Reader
	fileReader := bytes.NewReader(fileData)

	// 生成上传 token
	putPolicy := storage.PutPolicy{
		Scope: server.Bucket,
	}

	mac := qbox.NewMac(server.AccessKey, server.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	// 设置上传配置
	cfg := setConfig(server)

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, fileReader, fileSize, &putExtra)
	if err != nil {
		return "", err
	}

	url := "http://" + server.Domain + "/" + ret.Key
	return url, nil
}

func setConfig(server conf.OSS) *storage.Config {
	return &storage.Config{
		Region:        selectZone(server.Zone),
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
}

func selectZone(Zone int) *storage.Zone {
	switch Zone {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuadongZheJiang2
	case 3:
		return &storage.ZoneHuabei
	case 4:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
