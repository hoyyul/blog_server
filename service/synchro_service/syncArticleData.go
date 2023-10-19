package synchro_service

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/redis_service"
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func SyncArticleData() {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := redis_service.NewArticleDigg().GetInfo()
	lookInfo := redis_service.NewArticleVisit().GetInfo()
	commentInfo := redis_service.NewCommentCount().GetInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		json.Unmarshal(hit.Source, &article)

		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]

		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment
		if article.DiggCount == newDigg && article.LookCount == newLook && article.CommentCount == newComment {
			// no change
			continue
		}
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Infof("%s, like data synchronized successfully, number of likes %d, number of views %d, number of comments %d", article.Title, newDigg, newLook, newComment)

	}
	redis_service.NewArticleDigg().Clear()
	redis_service.NewArticleVisit().Clear()
	redis_service.NewCommentCount().Clear()
}
