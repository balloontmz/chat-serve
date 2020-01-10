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
func GroupList(uID int) []ChatGroup {
	var groups []ChatGroup
	var groupIDs []int
	DB.Model(&UserGroup{}).Where("user_id = ?", uID).Pluck("group_id", &groupIDs)
	DB.Model(&ChatGroup{}).Where("id in (?)", groupIDs).Find(&groups)
	return groups
}

//CreateGroup 创建聊天室
func CreateGroup(g ChatGroup) {
	DB.Create(&g)
	return
}

//GetGroupByID 根据 id 查询聊天室
func GetGroupByID(id string) ChatGroup {
	var g = ChatGroup{}
	DB.Where("id = ?", id).First(&g)
	return g
}
