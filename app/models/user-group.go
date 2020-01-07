package models

import (
	"github.com/labstack/gommon/log"
	"time"
)

//UserGroup 聊天信息对象
type UserGroup struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	UserID    string     `gorm:"type:int;index" json:"user_id"`
	GroupID   string     `gorm:"type:int;index" json:"group_id"`
}

//TableName 设置 UserGroup 的表名为`user_group`
func (UserGroup) TableName() string {
	return "chat_user_group"
}

//GetUsersUseGroupID 根据 group id 拉取用户
func GetUsersUseGroupID(groupID int) []User {
	return nil
}

//GetUserIDsByGroupID 根据聊天室 id 获取用户的 id
func GetUserIDsByGroupID(groupID int) []int {
	var ids []int
	DB.Model(&UserGroup{}).Where("group_id = ?", groupID).Pluck("user_id", &ids)
	log.Info("传入的 group id 为:", groupID)
	log.Info("根据 group id 查询到的用户为:", ids)
	return ids
}
