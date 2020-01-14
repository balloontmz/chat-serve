package chatmsg

import (
	"github.com/balloontmz/chat-serve/app/cusvalidate"
	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

//Index 获取聊天室列表
func Index(c echo.Context) error {

	var query cusvalidate.MsgListQuery

	log.Info("绑定数据之前")
	// 绑定数据
	if err := c.Bind(&query); err != nil {
		return res.Fmt(c, 0, "请求出错", err.Error())
	}
	log.Info("验证数据之前,此时的数据为:", query)
	// 验证数据
	if err := c.Validate(&query); err != nil {
		return res.Fmt(c, 0, "数据验证出错", err.Error())
	}
	return res.Fmt(c, 1, "", models.MsgList(query))
}
