package chatgroup

import (
	"net/http"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"
	"github.com/labstack/echo"
)

//Index 获取聊天室列表
func Index(c echo.Context) error {
	return res.Fmt(c, 1, "", models.GroupList())
}

//Store 保存聊天室
func Store(c echo.Context) error {
	g := models.ChatGroup{
		Name: c.FormValue("name"),
	}
	models.CreateGrouo(g)
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
