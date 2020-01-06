package chatmsg

import (
	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"

	"github.com/labstack/echo"
)

//Index 获取聊天室列表
func Index(c echo.Context) error {
	return res.Fmt(c, 1, "", models.MsgList())
}
