package models

import (
	"github.com/balloontmz/chat-serve/app/cusvalidate"
	"time"
)

//ChatMsg 聊天信息对象
type ChatMsg struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Msg       string     `gorm:"type:text" json:"msg"`
	Type      uint       `gorm:"type:int;index;default:1" json:"type"`
	UserID    uint       `gorm:"type:int;index" json:"user_id"`
	GroupID   uint       `gorm:"type:int;index" json:"group_id"`
	UserName  string     `gorm:"-" json:"user_name"`
}

//TableName 设置 ChatMsg 的表名为`chat_msg`
func (ChatMsg) TableName() string {
	return "chat_msg"
}

//MsgList 聊天信息列表
func MsgList(query cusvalidate.MsgListQuery) []ChatMsg {
	var d = DB.Table("chat_msg")
	if query.GroupIDS != nil {
		d = d.Where("group_id in (?)", query.GroupIDS)
	}

	var msgs []ChatMsg
	d.Select("chat_msg.id as id, chat_msg.msg as msg, chat_msg.type as type, chat_msg.user_id as user_id, chat_msg.group_id as group_id, chat_user.name as user_name").Joins("left join chat_user on chat_user.id = chat_msg.user_id").Scan(&msgs)
	// d.Find(&msgs)
	return msgs
}

//MsgListUseGroupIDs 根据传入的聊天室id 数组查找消息
func MsgListUseGroupIDs(groupIDs []int) []ChatMsg {
	var msgs []ChatMsg
	DB.Find(&msgs)
	return msgs
}

//CreateMsg 创建消息
func CreateMsg(userID, groupID, action int, msg string) ChatMsg {
	var m = ChatMsg{
		Msg:     msg,
		UserID:  uint(userID),
		GroupID: uint(groupID),
		Type:    uint(action),
	}
	DB.Create(&m)
	return m
}
