package config

import (
	"io"
	"os"
	"time"

	"github.com/labstack/gommon/log"
)

var (
	//GlobalDBLogger 全局的数据库日志
	GlobalDBLogger = log.New("pqsql")
)

//InitLog 初始化日志器的记录选项，初步测试成功
func InitLog() {
	// 重启程序时重新设置日志存放位置
	setUpGlobalLogger()

	setUpDBLogger() // 设置数据库日志器.

	go func() {

		for {
			// 每天零点执行一次日志生成命令
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)

			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

			t := time.NewTimer(next.Sub(now))

			<-t.C

			// 将日志写入文件，定时执行
			setUpGlobalLogger()

			setUpDBLogger()
		}
	}()
}

//设置全局日志器
func setUpGlobalLogger() {
	f, _ := os.OpenFile("log/echo"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // 追加或者新建文件

	w := io.MultiWriter(f)

	log.SetOutput(w)
	log.SetHeader(`[${time_rfc3339}] ${level}:`)
}

//设置数据库查询日志器
func setUpDBLogger() {
	if GlobalDBLogger != nil {
		// 将日志写入文件，定时执行
		f, _ := os.OpenFile("log/pqsql"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // 追加或者新建文件

		w := io.MultiWriter(f)

		GlobalDBLogger.SetOutput(w)
		GlobalDBLogger.SetHeader(`[${time_rfc3339}] ${level}:`)
	}

}
