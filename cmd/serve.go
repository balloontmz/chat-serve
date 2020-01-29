package cmd

import (
	"fmt"

	"github.com/balloontmz/chat-serve/app/config"
	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/routes"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run the serve",
		Long:  `run serve use tags like --port`,
		Run:   runServe,
	}
)

func init() {
	serveCmd.Flags().String("port", "1323", "serve port default 1323")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}

//runServe 启动网络服务
func runServe(cmd *cobra.Command, args []string) {
	fmt.Println("这里启动了服务器,当前获取的 port 为:", viper.GetString("port"))

	//初始化日志,每天自动创建存储日志
	config.InitLog() // 此配置初步测试成功，如果新建文件夹需先创建文件夹

	// 初始化数据库连接,可能需要添加连接池
	if _, err := models.InitDB(models.Config); err != nil {
		panic(err)
	}

	router := routes.NewEngine() // 初始化路由
	// router.GET("/ws", hello)
	// Listen and serve on 0.0.0.0:8080
	router.Logger.Fatal(router.Start(":8080"))
}
