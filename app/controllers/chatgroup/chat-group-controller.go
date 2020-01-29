package chatgroup

import (
	"net/http"
	"strconv"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"
	"github.com/balloontmz/chat-serve/app/service/wcservice"
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

//NotJoinGroup 获取用户未加入的聊天室列表
func NotJoinGroup(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtservice.JwtCustomClaims)
	return res.Fmt(c, 1, "", models.NotJoinGroupList(claims.UID))
}

//AddUser2Group 用户加入聊天组
func AddUser2Group(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtservice.JwtCustomClaims)
	gID, _ := strconv.Atoi(c.QueryParam("group_id"))
	uG := models.UserGroup{
		UserID:  uint(claims.UID),
		GroupID: uint(gID),
	}
	log.Info("将群组加入当前用户,群组 id 为:", gID, "用户 id 为:", claims)
	models.CreateUserGroup(&uG)
	return res.Fmt(c, 1, "添加群组成功", nil)
}

//Store 保存聊天室
func Store(c echo.Context) error {
	log.Info("当前请求的参数为: ", c.Request().Body, "当前请求的类型为:", c.Request().Header)
	g := models.ChatGroup{}
	if err := c.Bind(&g); err != nil {
		return res.Fmt(c, 0, "绑定数据失败", err)
	}

	log.Info("当前绑定完的数据为:", g)

	if e := models.GetGroupByName(g.Name); e.ID != 0 {
		log.Info("当前获取到的 group 为:", e)
		return res.Fmt(c, 0, "聊天室已存在", g)
	}
	models.CreateGroup(&g)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtservice.JwtCustomClaims)

	log.Debug("当前的用户 payload 为:", claims)

	uG := models.UserGroup{
		UserID:  uint(claims.UID),
		GroupID: g.ID,
	}
	models.CreateUserGroup(&uG)

	log.Info("当前的用户群聊关联为:", uG)

	return res.Fmt(c, 1, "添加成功", g)
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

//TestImage 测试图片解码
func TestImage(c echo.Context) error {
	str := c.FormValue("content")
	wcservice.Base64toBinary(str)
	return c.String(200, "test")
}
