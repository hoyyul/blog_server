package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/service/redis_service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	IsCollect bool `json:"is_collect"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var req models.ESIDRequest
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

	isCollect := IsUserArticleColl(c, model.ID)
	var articleDetail = ArticleDetailResponse{
		ArticleModel: model,
		IsCollect:    isCollect,
	}

	res.OkWithData(articleDetail, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func IsUserArticleColl(c *gin.Context, articleID string) (isCollect bool) {
	// if login
	token := c.GetHeader("token")
	if token == "" {
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		return
	}

	// if logout
	if redis_service.CheckLogout(token) {
		return
	}
	var count int64
	global.DB.Model(models.UserArticleCollectionModels{}).Where("user_id = ? and article_id =?", claims.UserID, articleID).Count(&count)
	if count == 0 {
		return
	}
	return true
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
