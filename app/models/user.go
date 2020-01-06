package models

import (
	"time"
)

//User 聊天信息对象
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Name      string     `gorm:"type:varchar(255);unique_index" json:"name"` // 唯一索引
	Password  string     `gorm:"type:varchar(255)" json:"-"`
}

//TableName 设置 User 的表名为`chat_msg`
func (User) TableName() string {
	return "chat_user"
}

//GetUserByUserName 根据用户名获取用户
func GetUserByUserName(name string) User {
	var u = User{}
	DB.Where("name = ?", name).First(&u)
	return u
}