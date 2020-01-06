package models

import (
	"time"
)

//ChatMsg 聊天信息对象
type ChatMsg struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Msg       string     `gorm:"type:text" json:"msg"`
	UserID    uint       `gorm:"type:int;index" json:"user_id"`
	GroupID   uint       `gorm:"type:int;index" json:"group_id"`
}

//TableName 设置 ChatMsg 的表名为`chat_msg`
func (ChatMsg) TableName() string {
	return "chat_msg"
}

//MsgList 聊天信息列表
func MsgList() []ChatMsg {
	var msgs []ChatMsg
	DB.Find(&msgs)
	return msgs
}
