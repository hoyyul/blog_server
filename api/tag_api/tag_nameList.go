package tag_api

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

type TagResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (TagApi) TagNameListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	query := elastic.NewBoolQuery()
	agg := elastic.NewTermsAggregation().Field("tags")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["tags"]
	var tagType T
	json.Unmarshal(byteData, &tagType)

	var tagList = make([]TagResponse, 0)
	for _, bucket := range tagType.Buckets {
		tagList = append(tagList, TagResponse{
			Label: bucket.Key,
			Value: bucket.Key,
		})
	}

	res.OkWithData(tagList, c)

}
