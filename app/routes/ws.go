package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/ws"

	"github.com/labstack/echo"
)

//WsRoutesRegister is a simple
func wsRoutesRegister(router *echo.Echo) *echo.Echo {

	var wsRouter = router.Group("ws")
	wsRouter.GET("", ws.Entrance)
	return router
}
