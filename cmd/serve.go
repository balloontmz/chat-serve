package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/balloontmz/chat-serve/app/config"
	"github.com/balloontmz/chat-serve/app/routes"
	"github.com/balloontmz/chat-serve/app/models"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
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

	//用于 websocket
	upgrader = websocket.Upgrader{}

	strChan = make(chan string, 100)

	activeMap [](chan string)
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
	router.GET("/ws", hello)
	// Listen and serve on 0.0.0.0:8080
	router.Logger.Fatal(router.Start(":1323"))
}

func hello(c echo.Context) error {
	fmt.Print("ws 请求进入此处")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // 不检查源
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	fmt.Print("查看是否有错误")
	if err != nil {
		fmt.Print("有错误", err)
		return err
	}
	fmt.Print("没有错误")
	defer ws.Close()

	var index int

	var activeChan = make(chan string, 10)

	go WriteMessage(ws, activeChan)

	activeMap = append(activeMap, activeChan)

	for {
		index++
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		for _, c := range activeMap {
			c <- (string(msg) + strconv.Itoa(index))
		}
		fmt.Printf("读取到来自客户端的消息,将其返回给客户端%s\n", msg)
	}
}

//WriteMessage 写入消息给客户端
func WriteMessage(ws *websocket.Conn, strChan <-chan string) {
	for {
		if str, ok := <-strChan; ok {
			err := ws.WriteMessage(websocket.TextMessage, []byte(str))
			// fmt.Print("发送给客户端字符串,为:", str)
			if err != nil {
				fmt.Print("当前写入的错误为:", err)
			}
		}
	}
}
