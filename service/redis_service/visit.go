package redis_service

import (
	"blog_server/global"
	"strconv"
)

const visitPrefix = "visit"

func Visit(id string) error {
	num, _ := global.Redis.HGet(visitPrefix, id).Int()
	num++
	err := global.Redis.HSet(visitPrefix, id, num).Err()
	return err
}

func GetVisit(id string) int {
	num, _ := global.Redis.HGet(visitPrefix, id).Int()
	return num
}

func GetVisitInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(visitPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func VisitClear() {
	global.Redis.Del(visitPrefix)
}
