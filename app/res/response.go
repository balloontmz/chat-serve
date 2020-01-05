package res

import (
	"net/http"

	"github.com/labstack/echo"
)

// ResponseFmt 返回格式结构体
type ResponseFmt struct {
	Ret  int         `json:"ret" xml:"ret"`
	Msg  string      `json:"msg" xml:"msg"`
	Data interface{} `json:"data" xml:"data"`
}

//Fmt 简化的返回值
// type Fmt struct {
// 	Ret  int         `json:"ret" xml:"ret"`
// 	Msg  string      `json:"msg" xml:"msg"`
// 	Data interface{} `json:"data" xml:"data"`
// }

//Fmt 格式化响应返回
func Fmt(c echo.Context, ret int, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, &ResponseFmt{ret, msg, data})
}

//ErrFmt 错误返回
func ErrFmt(c echo.Context, ret int, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, &ResponseFmt{ret, msg, data})
}
