//Package wcservice 生成词云图的服务
package wcservice

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/balloontmz/chat-serve/app/cusvalidate"
	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/service/qnservice"
	"github.com/labstack/gommon/log"
)

var (
	//WordCloudURL 云图生成网址
	WordCloudURL string = "http://0.0.0.0:10087"
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

//GetWordCloud 获取内容的云图
func GetWordCloud(content string) string {
	// params := url.Values{"content": []string{"aaa"}}
	params := url.Values{"content": []string{content}}
	resp, err := http.PostForm(WordCloudURL, params)
	if err != nil {
		log.Info("请求出错")
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))
	return string(body)
}

//UpdateGroupsWordCloud 更新群组的
func UpdateGroupsWordCloud() {
	groups := models.AllGroupList()
	for _, group := range groups {
		msgs := models.MsgList(cusvalidate.MsgListQuery{
			GroupIDS: []int{int(group.ID)},
		})
		content := ""
		for _, msg := range msgs {
			if msg.Type == 1 {
				content += msg.Msg
			}
		}
		//skip condition
		if content == "" {
			continue
		}

		base64Img := GetWordCloud(content)
		binaryImg := Base64toBinary(base64Img)
		url := qnservice.UploadByBytes(binaryImg)
		group.UpdateWordCloud(url)
	}
}
