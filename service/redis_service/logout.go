package redis_service

import (
	"blog_server/global"
	"time"
)

func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}
