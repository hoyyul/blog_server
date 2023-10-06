package es_service

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/redis_service"
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func GetList(option Option) (articleList []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()

	// search article by title, abstract, content
	if option.Key != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}

	// search article by tag
	if option.Tag != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	}

	// sort setting
	// default
	sortField := SortField{
		Field:     "created_at",
		Ascending: false,
	}
	if option.Sort != "" {
		_list := strings.Split(option.Sort, " ")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortField.Field = _list[0]
			if _list[1] == "desc" {
				sortField.Ascending = false
			}
			if _list[1] == "asc" {
				sortField.Ascending = true
			}
		}
	}

	//search
	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")). // highlight title if search by title
		From(option.GetForm()).
		Sort(sortField.Field, sortField.Ascending).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value)

	// save hit to struct
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel

		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}

		// hightlight title
		if title, ok := hit.Highlight["title"]; ok {
			article.Title = title[0]
		}

		article.ID = hit.Id // id = _id

		// get count from redis
		article.DiggCount = article.DiggCount + redis_service.NewArticleDigg().Get(hit.Id)
		article.LookCount = article.LookCount + redis_service.NewArticleVisit().Get(hit.Id)
		article.CommentCount = article.CommentCount + redis_service.NewCommentCount().Get(hit.Id)

		articleList = append(articleList, article)
	}
	return articleList, count, err
}

func GetDetail(id string) (article models.ArticleModel, err error) {
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		return
	}
	article.ID = res.Id

	// get count from redis
	article.DiggCount = article.DiggCount + redis_service.NewArticleDigg().Get(id)
	article.LookCount = article.LookCount + redis_service.NewArticleVisit().Get(id)
	article.CommentCount = article.CommentCount + redis_service.NewCommentCount().Get(id)
	return
}

func GetDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("article doesn't exist")
	}
	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return
}

func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).
		Refresh("true").
		Do(context.Background())
	return err
}
