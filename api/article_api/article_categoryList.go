package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type CategoryResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (ArticleApi) ArticleCategoryListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}

	agg := elastic.NewTermsAggregation().Field("category")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("categorys", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["categorys"]
	var categoryType T
	_ = json.Unmarshal(byteData, &categoryType)
	var categoryList = make([]CategoryResponse, 0)
	for _, bucket := range categoryType.Buckets {
		categoryList = append(categoryList, CategoryResponse{
			Label: bucket.Key,
			Value: bucket.Key,
		})
	}
	res.OkWithData(categoryList, c)

}
