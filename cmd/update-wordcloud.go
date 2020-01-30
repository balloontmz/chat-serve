package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/service/wcservice"
)

func init() {
	rootCmd.AddCommand(updateWordCloudCmd)
}

var (
	updateWordCloudCmd = &cobra.Command{
		Use:   "updateword",
		Short: "run the updateword",
		Long:  `run updateword use tags like`,
		Run:   runUpdateWordCloud,
	}
)

//runServe 启动网络服务
func runUpdateWordCloud(cmd *cobra.Command, args []string) {
	fmt.Print("开始更新")
	// 初始化数据库连接,可能需要添加连接池
	if _, err := models.InitDB(models.Config); err != nil {
		panic(err)
	}
	wcservice.UpdateGroupsWordCloud()
	fmt.Print("结束更新")
}
