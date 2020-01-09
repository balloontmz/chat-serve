package ws

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/balloontmz/chat-serve/app/res"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"
	"github.com/balloontmz/chat-serve/app/service/wsservice"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	//用于 websocket
	upGrader = websocket.Upgrader{}

	strChan = make(chan string, 100)

	activeMap [](chan string)
)

//Entrance websocket 入口 -- 一个用户一种客户端默认只有一个连接
func Entrance(c echo.Context) error {
	fmt.Print("ws 请求进入此处")
	log.Info("当前传入的 token 为:", c.QueryParam("token"))

	upGrader.CheckOrigin = func(r *http.Request) bool { return true } // 不检查源
	ws, err := upGrader.Upgrade(c.Response(), c.Request(), nil)

	//TODO: group id 应该是需要在发送的消息中携带, websocket 创建应该是进入列表时创建!!!现在的 msg 只是单纯的string,后面应该改造成对象!!!
	gID, _ := strconv.Atoi(c.QueryParam("g_id"))
	if gID == 0 {
		gID = 3
	}
	var uID = jwtservice.GetUserIDFromEchoContext(c)

	wsservice.SetConnMap(uID, ws) // 设置全局变量中用户连接的 map,如果需要分布式,需要对某个用户做一个映射

	fmt.Print("当前传入的 group id 为:", gID, "--------当前传入的 user id 为:", uID)
	if err != nil {
		fmt.Print("有错误", err)
		return err
	}
	fmt.Print("没有错误")
	defer ws.Close()

	// activeMap = append(activeMap, activeChan)

	var index int
	for {
		index++
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Info("进入此处代表读取消息失败,ws 连接中断", err)
			//此处将维护的全局 map 中移除该用户的连接
			return res.Fmt(c, 0, "ws 连接中断", nil)
		}
		//读取到一条消息之后
		//目前设想是: 将消息发送给 group 的处理器,保存到数据库并分发给当前活跃的用户 ws,存的应该有消息内容和发送者 id
		//对于发送者,应该是返回是否发送成功的标识,而不是消息内容.
		//消息是否应该存在一个唯一标识 -- id?
		wsservice.DealMsg(uID, gID, msg)

		fmt.Printf("读取到来自客户端的消息,将其返回给客户端%s\n", msg)
	}
}
