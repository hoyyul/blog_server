package jwts

import (
	"blog_server/global"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
)

// get claim with secret key and token
func ParseToken(tokenStr string) (*CustomClaim, error) {
	MySecret = []byte(global.Config.Jwt.SecretKey)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
