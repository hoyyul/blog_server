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

func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var req models.PageInfo
	_ = c.ShouldBindQuery(&req)

	boolQuery := elastic.NewBoolQuery()

	if req.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(req.Key, "title", "body"))
	}

	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		if body, ok := hit.Highlight["body"]; ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}

	res.OkWithList(fullTextList, count, c)
}
