package cusjwt

import "github.com/dgrijalva/jwt-go"

//JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var (
	//JWTSecret jwt secret
	JWTSecret = "jwt-secret"
)
