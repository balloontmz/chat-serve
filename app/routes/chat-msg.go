package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/chatmsg"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//chatMsgRoutesRegister is a simple
func chatMsgRoutesRegister(router *echo.Echo) *echo.Echo {
	var group = router.Group("msg")

	group.Use(middleware.JWTWithConfig(jwtservice.CreateJWTConfig())) // 权限中间件
	group.GET("", chatmsg.Index)
	return router
}
