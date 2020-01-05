package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/balloontmz/chat-serve/app/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(dbTestCmd) 
}

var (
	dbTestCmd = &cobra.Command{
		Use:   "dbtest",
		Short: "run the dbtest",
		Long:  `run dbtest use tags like`,
		Run:   runTest,
	}
)

func init() {
}

//runServe 启动网络服务
func runTest(cmd *cobra.Command, args []string) {	
	if _, err := model.InitDB(model.Config); err != nil { // 初始化数据库链接
		panic(err)
	}
}