package statistic_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumResponse struct {
	UserCount        int `json:"user_count"`
	ArticleCount     int `json:"article_count"`
	MessageCount     int `json:"message_count"`
	ChatGroupCount   int `json:"chat_group_count"`
	NowLoginCount    int `json:"now_login_count"`
	NowRegisterCount int `json:"now_sign_count"`
}

func (StatisticApi) StatisticSumView(c *gin.Context) {

	var userCount, articleCount, messageCount, ChatGroupCount int
	var nowLoginCount, nowRegisterCount int

	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = int(result.Hits.TotalHits.Value)
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	global.DB.Model(models.MessageModel{}).Select("count(id)").Scan(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&ChatGroupCount)
	global.DB.Model(models.LoginDataModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowRegisterCount)

	res.OkWithData(DataSumResponse{
		UserCount:        userCount,
		ArticleCount:     articleCount,
		MessageCount:     messageCount,
		ChatGroupCount:   ChatGroupCount,
		NowLoginCount:    nowLoginCount,
		NowRegisterCount: nowRegisterCount,
	}, c)
}
