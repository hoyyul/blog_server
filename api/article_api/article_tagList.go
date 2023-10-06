package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

func (ArticleApi) ArticleTagListView(c *gin.Context) {
	var req models.PageInfo
	_ = c.ShouldBindQuery(&req)

	// set default value
	if req.Limit == 0 {
		req.Limit = 10
	}
	offset := (req.Page - 1) * req.Limit
	if offset < 0 {
		offset = 0
	}

	// get total count CountAggregation
	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")). // NewCardinalityAggregation() no repeat
		Size(0).
		Do(context.Background())
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)

	// tags, keyword, page Aggregation
	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(req.Limit)) // paginate aggregation search by tags + title
	query := elastic.NewBoolQuery()

	result, _ = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())

	// save data to list
	var tagType TagsType
	var tagList = make([]*TagsResponse, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType) // save json to struct

	for _, bucket := range tagType.Buckets {
		// save article
		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		// save tag
		tagList = append(tagList, &TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount, // total count for a tag
			ArticleIDList: articleList,     // artical by title(keyword)
		})
		//tagStringList = append(tagStringList, bucket.Key)
	}

	/*var tagModelList []models.TagModel
	global.DB.Find(&tagModelList, "title in ?", tagStringList) // if empty tag invalid
	var tagDate = map[string]string{}
	for _, model := range tagModelList {
		tagDate[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	for _, response := range tagList {
		response.CreatedAt = tagDate[response.Tag]
	}*/

	res.OkWithList(tagList, count, c)
}
