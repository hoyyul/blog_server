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
	var req models.ESIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	redis_service.NewArticleVisit().Set(req.ID) // visit

	result, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(req.ID).
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
