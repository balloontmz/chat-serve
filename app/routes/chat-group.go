package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers"
	"github.com/labstack/echo"
)

//EmailRoutesRegister is a simple
func chatGrouopRoutesRegister(router *echo.Echo) *echo.Echo {

	var group = router.Group("group")
	group.GET("", controllers.Index)
	group.POST("", controllers.Store)
	group.GET(":id", controllers.Show)
	// group.PUT("/:id", controllers.Update)
	// group.DELETE("/:id", controllers.Destroy)
	return router
}
