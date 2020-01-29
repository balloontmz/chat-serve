package qiniu

// 存储相关功能的引入包只有这两个，后面不再赘述
import (
	"github.com/balloontmz/chat-serve/app/res"
	"github.com/balloontmz/chat-serve/app/service/qnservice"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

//GetUploadToken 1
func GetUploadToken(c echo.Context) error {
	log.Info("进入此处")
	// accessKey := "NM-OrGXyrsOYB-zt5UqEsU7uYjWGgB4DWttRQz2Z"
	// secretKey := "mRk44qo5oHsvHHJy1QbqxVWh7yS1sU8xLL_cblk9"
	// bucket := "tomtiddler"
	accessKey := qnservice.AccessKey
	secretKey := qnservice.SecretKey
	bucket := qnservice.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)

	upToken := putPolicy.UploadToken(mac)
	log.Info("当前返回给客户端的 token 为:", upToken)
	return res.Fmt(c, 1, "", upToken)
}
