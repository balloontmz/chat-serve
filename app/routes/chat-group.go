package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/chatgroup"
	"github.com/labstack/echo"
)

//chatGroupRoutesRegister is a simple
func chatGroupRoutesRegister(router *echo.Echo) *echo.Echo {

	var group = router.Group("group")
	group.GET("", chatgroup.Index)
	group.POST("", chatgroup.Store)
	group.GET(":id", chatgroup.Show)
	// group.PUT("/:id", chatgroup.Update)
	// group.DELETE("/:id", chatgroup.Destroy)
	return router
}
