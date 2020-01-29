package qnservice

import (
	"bytes"
	"context"
	"fmt"

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
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// putExtra := storage.PutExtra{
	// 	Params: map[string]string{
	// 		"x:name": "github logo",
	// 	},
	// }
	data := b
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, "", bytes.NewReader(data), dataLen, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(ret.Key, ret.Hash)
	return ret.Key
}
