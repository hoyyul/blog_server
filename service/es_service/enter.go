package es_service

import (
	"blog_server/global"
	"blog_server/models"
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func GetList(key string, page, limit int) (articleList []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	from := page

	// search key from title
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}

	// set default value
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	//search
	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value)

	// save hit to struct
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel

		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}

		err = json.Unmarshal(data, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}

		article.ID = hit.Id
		articleList = append(articleList, article)
	}
	return articleList, count, err
}

func GetDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	return
}
