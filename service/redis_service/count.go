package redis_service

import (
	"blog_server/global"
	"strconv"
)

type CountDB struct {
	Index string // index
}

func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	//fmt.Println(num)
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()
	return err
}

func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, id, newNum).Err()
	return err
}

func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
