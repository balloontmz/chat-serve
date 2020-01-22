package routes

import (
	"github.com/balloontmz/chat-serve/app/controllers/qiniu"
	"github.com/labstack/echo"
)

//qiniuRoutesRegister 文档
func qiniuRoutesRegister(router *echo.Echo) *echo.Echo {

	var q = router.Group("upload")
	// uploadRouter.GET("", qiniu.GetUploadToken)

	// uploadRouter.Use(middleware.JWTWithConfig(jwtservice.CreateJWTConfig())) // 权限中间件
	q.GET("", qiniu.GetUploadToken)

	return router
}
