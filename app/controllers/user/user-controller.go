package user

import (
	"net/http"
	"time"

	"github.com/balloontmz/chat-serve/app/models"
	"github.com/balloontmz/chat-serve/app/res"
	"github.com/balloontmz/chat-serve/app/service/jwtservice"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//User 用户数据结构
type User struct {
	UserName string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

//Register 用户注册
func Register(c echo.Context) error {
	return nil
}

//Login 用户登录
func Login(c echo.Context) error {

	var uData = &User{}

	if err := c.Bind(uData); err != nil {
		return res.ErrFmt(c, 0, "未找到用户", nil)
	}

	var u = models.GetUserByUserName(uData.UserName)

	if u.Name == "" {
		return res.ErrFmt(c, 0, "未找到用户", nil)
	}

	// Throws unauthorized error
	if uData.UserName != u.Name || uData.Password != u.Password {
		return echo.ErrUnauthorized
	}

	claims := &jwtservice.JwtCustomClaims{
		Name: u.Name,
		UID:  int(u.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtservice.JWTSecret))
	if err != nil {
		return err
	}

	return res.Fmt(c, 1, "", echo.Map{
		"access_token": t,
		"token_type":   "Bearer",
		"username":     u.Name,
		"userid":       u.ID,
	})
}

//Info 用户信息
func Info(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtservice.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
