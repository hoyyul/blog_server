package redis_service

import (
	"blog_server/global"
	"encoding/json"
	"fmt"
	"time"
)

const newsIndex = "news_index"

type NewsData struct {
	Index    interface{} `json:"index"`
	Title    string      `json:"title"`
	HotValue string      `json:"hotValue"`
	Link     string      `json:"link"`
}

func SetNews(key string, newsDatas []NewsData) error {
	byteData, _ := json.Marshal(newsDatas)
	err := global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 1*time.Hour).Err()
	return err
}

func GetNews(key string) (newsDatas []NewsData, err error) {
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newsDatas)
	return
}
