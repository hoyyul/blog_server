package jwts

import (
	"blog_server/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// generate a token with secret key
func GenerateToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.SecretKey)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // expire time
			Issuer:    global.Config.Jwt.Issuer,                                                     // issuer
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
