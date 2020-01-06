package models

import (
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
