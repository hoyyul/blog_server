package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCollectView(c *gin.Context) {
	var req models.ESIDRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	model, err := es_service.GetDetail(req.ID)
	if err != nil {
		res.FailWithMessage("Article doesn't exist", c)
		return
	}

	var collection models.UserArticleCollectionModels
	err = global.DB.Take(&collection, "user_id = ? and article_id = ?", claim.UserID, req.ID).Error

	if err != nil {
		// Article not found in collection, so add it.
		err = global.DB.Create(&models.UserArticleCollectionModels{
			UserID:    claim.UserID,
			ArticleID: req.ID,
		}).Error
		if err != nil {
			// Handle error during creation
			res.FailWithMessage("Failed to collect the article", c)
			return
		}
		es_service.ArticleUpdate(req.ID, map[string]any{
			"collects_count": model.CollectsCount + 1,
		})
		res.OkWithMessage("Collect article successfully!", c)
		return
	}

	// If article found in collection, remove it.
	err = global.DB.Delete(&collection).Error
	if err != nil {
		// Handle error during deletion
		res.FailWithMessage("Failed to suspend collection", c)
		return
	}
	es_service.ArticleUpdate(req.ID, map[string]any{
		"collects_count": model.CollectsCount - 1,
	})
	res.OkWithMessage("Suspend collection!", c)
}
