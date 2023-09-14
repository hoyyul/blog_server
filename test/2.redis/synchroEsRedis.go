package main

import (
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/models"
	"blog_server/service/redis_service"
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	initialization.InitConf()
	global.Logger = initialization.InitLogger()
	global.Redis = initialization.ConnectRedisDB(0)
	global.ESClient = initialization.EsConnect()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := redis_service.GetDiggInfo()
	visitInfo := redis_service.GetVisitInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel

		// get dig count from es and redis
		json.Unmarshal(hit.Source, &article) // dig count from es
		digg := diggInfo[hit.Id]             // dig count from redis
		visit := visitInfo[hit.Id]
		newDigg := article.DiggCount + digg // redis + es
		newVisit := article.LookCount + visit

		if article.DiggCount == newDigg && article.LookCount == newVisit {
			continue
		}

		// synchronize dig count to es
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
				"look_count": newVisit,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Info(article.Title, "Synchronize dig count and visit count successfully", newDigg)
	}

	// clear cache
	redis_service.DiggClear()
	redis_service.VisitClear()

}
