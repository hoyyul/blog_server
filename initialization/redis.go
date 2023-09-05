package initialization

import (
	"blog_server/global"
	"context"
	"time"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       0,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Logger.Errorf("Failed to connect redis %s", redisConf.Addr())
		return nil
	}
	return rdb
}
