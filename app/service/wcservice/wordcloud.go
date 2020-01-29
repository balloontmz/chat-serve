//Package wcservice 生成词云图的服务
package wcservice

import (
	"encoding/base64"
	"io/ioutil"
	"net/url"
)

//Base64toBinary base64 转二进制
func Base64toBinary(bStr string) []byte {
	bStr, _ = url.QueryUnescape(bStr) // 先 url decode -- 和 python quote_plus 一致 https://stackoverflow.com/questions/27556773/urllib-quote-in-go
	var prefix = "data:image/png;base64,"
	// log.Info("查看一下 byte string 的全部", bStr)
	bStr = bStr[len(prefix):]
	// log.Info("再查看一下 byte string 的全部", bStr)

	imgByte, _ := base64.StdEncoding.DecodeString(bStr)
	ioutil.WriteFile("./test.png", imgByte, 0666) // 写入文件测试
	return imgByte
}
