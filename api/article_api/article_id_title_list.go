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

type ArticleIDTitleListResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (ArticleApi) ArticleIDTitleListView(c *gin.Context) {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Source(`{"_source": ["title"]}`).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to search", c)
		return
	}
	var articleIDTitleList = make([]ArticleIDTitleListResponse, 0)
	for _, hit := range result.Hits.Hits {
		var model models.ArticleModel
		json.Unmarshal(hit.Source, &model)
		articleIDTitleList = append(articleIDTitleList, ArticleIDTitleListResponse{
			Value: hit.Id,
			Label: model.Title,
		})
	}

	res.OkWithData(articleIDTitleList, c)

}
