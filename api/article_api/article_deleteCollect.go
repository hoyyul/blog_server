package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/utils/jwts"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) ArticleDeleteCollectView(c *gin.Context) {
	var req models.ESIDListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	var collects []models.UserArticleCollectionModels
	var articleIDList []string
	global.DB.Find(&collects, "user_id = ? and article_id in ?", claim.UserID, req.IDList).
		Select("article_id").
		Scan(&articleIDList)
	if len(articleIDList) == 0 {
		res.FailWithMessage("Illegal request", c)
		return
	}

	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}

	// update collect count in es
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		count := article.CollectsCount - 1
		err = es_service.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			global.Logger.Error(err)
			continue
		}
	}
	global.DB.Delete(&collects)
	res.OkWithMessage(fmt.Sprintf("Cancel collect %d articles", len(idList)), c)

}
