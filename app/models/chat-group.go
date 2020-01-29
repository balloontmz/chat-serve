package models

import (
	"github.com/labstack/gommon/log"
	"time"
)

//ChatGroup 聊天室对象
type ChatGroup struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name      string     `gorm:"size:255" json:"name"`                  // string默认长度为255, 使用这种tag重设。
	Avatar    string     `gorm:"size:255;default:''" json:"avatar"`     // string默认长度为255, 使用这种tag重设。
	WordCloud string     `gorm:"size:255;default:''" json:"word_cloud"` // string默认长度为255, 使用这种tag重设。
}

//Params 获取结果的参数
type Params struct {
	Page int
	Size int
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

//AllGroupList 列表
//TODO: 所有聊天室,后期把重复的列表放入一个方法!!!
func AllGroupList() []ChatGroup {
	var groups []ChatGroup
	DB.Model(&ChatGroup{}).Find(&groups)
	return groups
}

//NotJoinGroupList 不属于当前用户的聊天室列表
func NotJoinGroupList(uID int) []ChatGroup {
	var groups []ChatGroup
	var groupIDs []int
	DB.Model(&UserGroup{}).Where("user_id = ?", uID).Pluck("group_id", &groupIDs)
	DB.Model(&ChatGroup{}).Where("id not in (?)", groupIDs).Find(&groups)
	return groups
}

//GetGroups 根据请求的参数获取商品
func GetGroups(p Params) []ChatGroup {
	page := p.Page
	size := p.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	var chatGroup []ChatGroup
	DB.Offset((page - 1) * size).Limit(size).Find(&chatGroup)
	return chatGroup
}

//CreateGroup 创建聊天室
func CreateGroup(g *ChatGroup) {
	log.Info("当前需要创建的群组值为:", *g, "当点 newrecord 为:", DB.NewRecord(*g))
	DB.Create(g)
	log.Info("当点 newrecord 为:", DB.NewRecord(*g))
	return
}

//GetGroupByID 根据 id 查询聊天室
func GetGroupByID(id string) ChatGroup {
	var g = ChatGroup{}
	DB.Where("id = ?", id).First(&g)
	return g
}

//GetGroupByName 根据名字查找聊天室
func GetGroupByName(name string) ChatGroup {
	var g = ChatGroup{}
	DB.Where("name = ?", name).First(&g)
	return g
}

//UpdateWordCloud 更新群组的云图
func (g ChatGroup) UpdateWordCloud(url string) {
	g.WordCloud = url
	DB.Save(&g)
}
