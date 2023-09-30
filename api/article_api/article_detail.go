package article_api

import (
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/service/redis_service"

	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var req ESIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	redis_service.NewArticleVisit().Set(req.ID) // visit
	model, err := es_service.GetDetail(req.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var req ArticleDetailRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	model, err := es_service.GetDetailByKeyword(req.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}