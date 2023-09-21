package redis_service

import (
	"blog_server/global"
	"strconv"
)

const commentPrefix = "comment"

func Comment(id string) error {
	num, _ := global.Redis.HGet(commentPrefix, id).Int()
	num++
	err := global.Redis.HSet(commentPrefix, id, num).Err()
	return err
}

func GetComment(id string) int {
	num, _ := global.Redis.HGet(commentPrefix, id).Int()
	return num
}

func GetCommentInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(commentPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func CommentClear() {
	global.Redis.Del(commentPrefix)
}
