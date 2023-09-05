package redis_service

import (
	"blog_server/global"
	"blog_server/utils"
)

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	return utils.InList(prefix+token, keys)
}
