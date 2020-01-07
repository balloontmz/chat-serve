package jwtservice

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	//JwtCustomClaims are custom claims extending default ones.
	JwtCustomClaims struct {
		Name string `json:"name"`
		UID  int    `json:"u_id"`
		jwt.StandardClaims
	}
)

var (
	//JWTSecret jwt secret
	JWTSecret = "jwt-secret"
)

func init() {

}

//Valid 实现 Claims interface,对原 Valid 方法进行扩展
func (j JwtCustomClaims) Valid() error {
	if err := j.StandardClaims.Valid(); err != nil {
		return err
	}
	//用户登录完成应该在缓存中存储一个和过期时间等同或稍大的值,用于此处拉取并对比
	// cacheClaim = getCache(j.Name)
	// if cacheClaim.Id != j.Id {
	// 	return errors.New("another login") // 别处已登录,此处的 jwt 不合法
	// }
	return nil
}

//CreateJWTConfig 创建一个新的 jwt config 对象 -- 此对象不可设置为全局对象,因为每个请求该对象的值是不一样的
func CreateJWTConfig() middleware.JWTConfig {
	var c = middleware.JWTConfig{
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
		Claims:        &JwtCustomClaims{},
		SigningKey:    []byte(JWTSecret),
		// ErrorHandler: ErrorHandler,     // 定义验证失败的返回
		// SuccessHandler: SuccessHandler,
	}
	return c
}

//CreateJWTConfigFromQuery 创建一个新的 jwt config 对象 -- 此对象不可设置为全局对象,因为每个请求该对象的值是不一样的
func CreateJWTConfigFromQuery() middleware.JWTConfig {
	var c = middleware.JWTConfig{
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "user",
		TokenLookup:   "query:token",
		AuthScheme:    "Bearer",
		Claims:        &JwtCustomClaims{},
		SigningKey:    []byte(JWTSecret),
		// ErrorHandler: ErrorHandler,     // 定义验证失败的返回
		// SuccessHandler: SuccessHandler,
	}
	return c
}

//GetUserIDFromEchoContext 从 echo 上下文当中获取当前用户的 id
func GetUserIDFromEchoContext(c echo.Context) int {
	var user = c.Get("user").(*jwt.Token)
	var claims = user.Claims.(*JwtCustomClaims)
	return claims.UID
}
