package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/redis_service"
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleContentByIDView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	redis_service.NewArticleVisit().Set(cr.ID)

	result, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage("Failed to search", c)
		return
	}
	var model models.ArticleModel
	err = json.Unmarshal(result.Source, &model)
	if err != nil {
		return
	}
	res.OkWithData(model.Content, c)
}
