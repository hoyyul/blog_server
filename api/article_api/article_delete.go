package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}

	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to delete article", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("Deleted %d articles successfully", len(result.Succeeded())), c)

}