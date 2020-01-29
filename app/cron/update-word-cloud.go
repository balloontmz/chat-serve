package config

import (
	"time"

	"github.com/balloontmz/chat-serve/app/service/wcservice"
)

//Run 运行每日任务
func Run() {
	go func() {

		for {
			// 每天零点执行一次日志生成命令
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)

			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

			t := time.NewTimer(next.Sub(now))

			<-t.C

			updateWordCloud()
		}
	}()
}

//每日零时更新云图
func updateWordCloud() {
	wcservice.UpdateGroupsWordCloud()
}
