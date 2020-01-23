package chatgroup

import (
	"net/http"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//Index 获取聊天室列表
func Index(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtservice.JwtCustomClaims)
	return res.Fmt(c, 1, "", models.GroupList(claims.UID))
}

//Store 保存聊天室
func Store(c echo.Context) error {
	log.Info("当前请求的参数为: ", c.Request().Body, "当前请求的类型为:", c.Request().Header)
	g := models.ChatGroup{
		Name:   c.FormValue("name"),
		Avatar: c.FormValue("avatar"),
	}
	if e := models.GetGroupByName(g.Name); e.ID != 0 {
		log.Info("当前获取到的 group 为:", e)
		return res.Fmt(c, 0, "聊天室已存在", g)
	}
	models.CreateGroup(&g)
	return res.Fmt(c, 1, "", g)
}

//Show 显示聊天室
func Show(c echo.Context) error {
	id := c.Param("id")
	return res.Fmt(c, 1, "", models.GetGroupByID(id))
}

//Update 更新聊天室
func Update(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

//Destroy 删除聊天室
func Destroy(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
