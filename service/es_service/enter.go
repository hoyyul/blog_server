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

type Option struct {
	models.PageInfo          // pagnation info
	Fields          []string // search article by title, abstract, content
	Tag             string   // search article by tag
}

type SortField struct {
	Field     string
	Ascending bool
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}

func GetList(option Option) (articleList []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()

	// search article by title, abstract, content
	if option.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}

	// search article by tag
	if option.Tag != "" {
		boolSearch.Must(
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
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value)

	diggInfo := redis_service.GetDiggInfo()
	visitInfo := redis_service.GetVisitInfo()
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

		// hightlight title
		if title, ok := hit.Highlight["title"]; ok {
			article.Title = title[0]
		}

		article.ID = hit.Id // id = _id

		// get dig count from redis
		digg := diggInfo[hit.Id]
		visit := visitInfo[hit.Id]
		article.DiggCount = article.DiggCount + digg
		article.LookCount = article.LookCount + visit

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
	article.LookCount = article.LookCount + redis_service.GetVisit(res.Id)
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
		Do(context.Background())
	return err
}
