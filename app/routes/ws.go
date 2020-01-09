package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/ws"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//WsRoutesRegister is a simple
func wsRoutesRegister(router *echo.Echo) *echo.Echo {

	var wsRouter = router.Group("ws")
	wsRouter.Use(middleware.JWTWithConfig(jwtservice.CreateJWTConfig())) // 权限中间件
	wsRouter.GET("", ws.Entrance)
	return router
}
