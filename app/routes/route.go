package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	// "gopkg.in/go-playground/validator.v9"
)

//NewEngine is a simple
func NewEngine() *echo.Echo {

	// 创建一个不带任何默认中间件的路由器
	router := echo.New()

	// router.Validator = &cusvalidator.CustomValidator{Validator: validator.New()}

	// 全局拦截器
	// 即使您使用GIN_MODE = release进行设置，Logger中间件也会将日志写入gin.DefaultWriter。
	// log的默认输出为 gin.DefaultWriter = os.Stdout
	router.Use(middleware.Logger())

	// 恢复中间件从任何错误中恢复，如果有错误的话，写入500。
	router.Use(middleware.Recover())

	router.Use(middleware.CORS())

	// 以下为路由,其中可能包括群组中间件和单中间件
	//--------------------------------------------------------------------------------------------------------------------
	router.GET("/", func(c echo.Context) error {
		log.Info("测试首页用")
		return c.String(200, "pong")
	})

	wsRoutesRegister(router)
	chatGroupRoutesRegister(router)
	userRoutesRegister(router)
	//---------------------------------------------------------------------------------------------------------------------

	// 静态文件路由
	router.Static("/apidoc", "./apidoc")

	return router
}
