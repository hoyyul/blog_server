package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// payload in custom claim
type JwtPayLoad struct {
	//Username string `json:"username"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"` // 1.admin 2. user
	UserID   uint   `json:"user_id"`
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
