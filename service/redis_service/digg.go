package redis_service

import (
	"blog_server/global"
	"strconv"
)

const diggPrefix = "digg"

func Digg(id string) error {
	num, _ := global.Redis.HGet(diggPrefix, id).Int() // hash table "digg", hash key id
	num++
	err := global.Redis.HSet(diggPrefix, id, num).Err()
	return err
}

func GetDiggCount(id string) int {
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	return num
}

// get article - digg count map
func GetDiggInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

// clear cache
func DiggClear() {
	global.Redis.Del(diggPrefix)
}
