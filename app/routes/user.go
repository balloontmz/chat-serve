package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/user"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//userRoutesRegister 用户相关路由注册
func userRoutesRegister(router *echo.Echo) *echo.Echo {
	var u = router.Group("user")
	u.POST("/login", user.Login)

	// Configure middleware with the custom claims type
	u.Use(middleware.JWTWithConfig(jwtservice.CreateJWTConfig())) // 权限中间件

	u.GET("/info", user.Info)

	return router
}
