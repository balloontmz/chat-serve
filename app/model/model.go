package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // 数据库驱动
)

type (
	// DBConfig 数据库配置结构体
	DBConfig struct {
		Username string `ini:"username"`
		Password string `ini:"password"`
		Hostname string `ini:"hostname"`
		Port     string `ini:"port"`
		Database string `ini:"database"`
	}
)

var (
	// DB 数据库引擎
	DB *gorm.DB
	// Config 数据库配置
	Config *DBConfig
)

func init() {
	Config = &DBConfig{
		Username: "chat",
		Password: "chat",
		Hostname: "47.100.124.234",
		Port:     "5432",
		Database: "chat",
	}
}

// InitDB 初始化数据库引擎
func InitDB(conf *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			conf.Hostname, conf.Port, conf.Username, conf.Database, conf.Password))
	if err != nil {
		return nil, err
	}
	DB = db
	Config = conf

	go func(db *gorm.DB) {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for {
			<-ticker.C
			db.DB().Ping()
		}
	}(db)
	return db, err
}

// GetDB 获取数据库连接引擎
func GetDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	if nil == Config {
		err := fmt.Errorf("no db config")
		return nil, err
	}

	// reconnect
	return InitDB(Config)
}
