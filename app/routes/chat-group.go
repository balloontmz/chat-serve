package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/chatgroup"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//chatGroupRoutesRegister is a simple
func chatGroupRoutesRegister(router *echo.Echo) *echo.Echo {
	var group = router.Group("group")

	group.Use(middleware.JWTWithConfig(jwtservice.CreateJWTConfig())) // 权限中间件
	group.GET("/not-join", chatgroup.NotJoinGroup)
	group.GET("/find-list", chatgroup.FindList)
	group.GET("", chatgroup.Index)
	group.POST("", chatgroup.Store)
	group.POST("/add-user", chatgroup.AddUser2Group)
	group.GET(":id", chatgroup.Show)
	// group.PUT("/:id", chatgroup.Update)
	// group.DELETE("/:id", chatgroup.Destroy)
	return router
}
