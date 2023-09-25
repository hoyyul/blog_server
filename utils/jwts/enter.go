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
	Avatar   string `json:"avatar"`
}

var MySecret []byte

type CustomClaim struct {
	JwtPayLoad
	jwt.StandardClaims
}
