package qnservice

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	//AccessKey 访问 key
	AccessKey string = "NM-OrGXyrsOYB-zt5UqEsU7uYjWGgB4DWttRQz2Z"
	//SecretKey 访问 key
	SecretKey string = "mRk44qo5oHsvHHJy1QbqxVWh7yS1sU8xLL_cblk9"
	//Bucket 访问 key
	Bucket string = "tomtiddler"

	prefix string = "http://qiniu.tomtiddler.top/"
)

//UploadByBytes 服务端上传字节数组到空间
func UploadByBytes(b []byte) string {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	// cfg.Zone = &storage.ZoneHuanan
	reg, _ := storage.GetRegionByID(storage.RIDHuanan)
	cfg.Region = &reg
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	data := b
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, calHash(b), bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(ret.Key, ret.Hash)
	return prefix + ret.Key
}

func calHash(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	s := base64.StdEncoding.EncodeToString(cipherStr)
	log.Info("计算出来的文件名为:", s)

	return s
}
