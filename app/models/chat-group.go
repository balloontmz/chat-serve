package models

import (
	"time"
)

//ChatGroup 聊天室对象
type ChatGroup struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `gorm:"size:255" json:"name"` // string默认长度为255, 使用这种tag重设。
}

//TableName 设置 ChatGroup 的表名为`chat_group`
func (ChatGroup) TableName() string {
	return "chat_group"
}

//GroupList 聊天室列表
func GroupList() []ChatGroup {
	var groups []ChatGroup
	DB.Find(&groups)
	return groups
}

//CreateGrouo 创建聊天室
func CreateGrouo(g ChatGroup) {
	DB.Create(&g)
	return
}

//GetGroupByID 根据 id 查询聊天室
func GetGroupByID(id string) ChatGroup {
	var g = ChatGroup{}
	DB.Where("id = ?", id).First(&g)
	return g
}
