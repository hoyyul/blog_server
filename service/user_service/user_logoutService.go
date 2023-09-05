package user_service

import (
	"blog_server/service/redis_service"
	"blog_server/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaim, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_service.Logout(token, diff)
}
