package wsservice

import (
	"fmt"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
)

var (
	//GlobalUserWS 全局的 user socket 连接 map
	GlobalUserWS map[int]*websocket.Conn
)

func init() {
	//先初始化,可能不必要初始化
	GlobalUserWS = make(map[int]*websocket.Conn)
}

//DealMsg 处理 ws 发送过来的消息
func DealMsg(userID, groupID int, msg string) {
	var uIDs = getUsersUseGroupID(groupID) // 获取当前聊天室的所有用户id
	go sendMsgUseIDs(uIDs, msg)            // 此处放入异步处理 -- 此处应该放入队列,并且有一定的除错机制
	insertMsg(userID, groupID, msg)        // 将消息加入数据库
}

//SetConnMap 设置全局 wsConn map 中某个用户的值
func SetConnMap(uID int, ws *websocket.Conn) {
	//此处应该考虑之前存在 conn 的情况
	GlobalUserWS[uID] = ws
}

func getUsersUseGroupID(groupID int) []int {
	return models.GetUserIDsByGroupID(groupID)
}

func insertMsg(uID, gID int, msg string) {
	return
}

func sendMsgUseIDs(uIDs []int, msg string) {
	if len(uIDs) == 0 {
		log.Info("尝试返送消息时没有查询到相关用户")
	}
	for _, id := range uIDs {
		//查找全局变量中是否存在当前用户的 ws conn.如果有,就发送消息
		if c, ok := GlobalUserWS[id]; ok {
			err := c.WriteMessage(websocket.TextMessage, []byte(msg))
			// fmt.Print("发送给客户端字符串,为:", str)
			if err != nil {
				fmt.Print("当前发送消息给用户,写入的错误为:", err, "用户 id 为:", id) // 记录错误,可能此处需要从全局中删除该变量!!!
			}
		}
	}
}
