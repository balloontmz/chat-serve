package wsservice

import (
	"encoding/json"
	"fmt"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
)

var (
	//GlobalUserWS 全局的 user socket 连接 map
	GlobalUserWS map[int]*websocket.Conn
)

type (
	//WSMsgStruct ws 消息结构体
	WSMsgStruct struct {
		Action int            `json:"action"`
		Data   GroupMsgStruct `json:"data"`
	}

	//GroupMsgStruct 聊天室消息结构体
	GroupMsgStruct struct {
		GroupID int    `json:"group_id"`
		Msg     string `json:"msg"`
	}
)

func init() {
	//先初始化,可能不必要初始化
	GlobalUserWS = make(map[int]*websocket.Conn)
}

//DealMsg 处理 ws 发送过来的消息
func DealMsg(userID, groupID int, msg []byte) error {
	var msgStruct WSMsgStruct

	err := json.Unmarshal(msg, &msgStruct)

	if err != nil {
		return err
	}

	switch msgStruct.Action {
	case 1, 2: // 代表是聊天室消息和图片

		if err != nil {
			return err
		}
		var uIDs = getUsersUseGroupID(msgStruct.Data.GroupID)                                          // 获取当前聊天室的所有用户id
		var msgModel = insertMsg(userID, msgStruct.Data.GroupID, msgStruct.Action, msgStruct.Data.Msg) // 将消息加入数据库
		go sendMsgUseIDs(uIDs, msgModel)                                                               // 此处放入异步处理 -- 此处应该放入队列,并且有一定的除错机制
	default:
		var uIDs = getUsersUseGroupID(msgStruct.Data.GroupID)                                          // 获取当前聊天室的所有用户id
		var msgModel = insertMsg(userID, msgStruct.Data.GroupID, msgStruct.Action, msgStruct.Data.Msg) // 将消息加入数据库
		go sendMsgUseIDs(uIDs, msgModel)                                                               // 此处放入异步处理 -- 此处应该放入队列,并且有一定的除错机制
	}

	return nil
}

//SetConnMap 设置全局 wsConn map 中某个用户的值
func SetConnMap(uID int, ws *websocket.Conn) {
	//此处应该考虑之前存在 conn 的情况
	GlobalUserWS[uID] = ws
}

func getUsersUseGroupID(groupID int) []int {
	return models.GetUserIDsByGroupID(groupID)
}

func insertMsg(uID, gID, action int, msg string) models.ChatMsg {
	return models.CreateMsg(uID, gID, action, msg)
}

func sendMsgUseIDs(uIDs []int, msgModel models.ChatMsg) {
	if len(uIDs) == 0 {
		log.Info("尝试返送消息时没有查询到相关用户")
	}
	u := models.GetUserByUserID(msgModel.UserID)
	msgModel.UserName = u.Name
	msg, err := json.Marshal(msgModel)
	if err != nil {
		log.Info("序列化消息失败,原因为:", err)
		return
	}

	for _, id := range uIDs {
		//查找全局变量中是否存在当前用户的 ws conn.如果有,就发送消息
		log.Info("发送消息中,用户 id 为:", id)
		if c, ok := GlobalUserWS[id]; ok {

			err := c.WriteMessage(websocket.TextMessage, []byte(msg))
			// fmt.Print("发送给客户端字符串,为:", str)
			if err != nil {
				fmt.Print("当前发送消息给用户,写入的错误为:", err, "用户 id 为:", id) // 记录错误,可能此处需要从全局中删除该变量!!!
			}
		} else {
			log.Info("进入此处代表没查找到当前用户的 ws conn")
		}
	}
}
