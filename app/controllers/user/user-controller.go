package user

import (
	"net/http"
	"time"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"
	"github.com/balloontmz/chat-serve/app/service/cusjwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//Register 用户注册
func Register(c echo.Context) error {
	return nil
}

//Login 用户登录
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var u = models.GetUserByUserName(username)

	if u.Name == "" {
		return res.ErrFmt(c, 0, "未找到用户", nil)
	}

	// Throws unauthorized error
	if username != u.Name || password != u.Password {
		return echo.ErrUnauthorized
	}

	claims := &cusjwt.JwtCustomClaims{
		Name: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(cusjwt.JWTSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

//Info 用户信息
func Info(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cusjwt.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
