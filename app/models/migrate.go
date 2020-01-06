package models

// Migrate 数据库迁移
func Migrate() {
	var db, _ = GetDB()
	db.AutoMigrate(&ChatGroup{}, &ChatMsg{}, &UserGroup{}, &User{})
}
