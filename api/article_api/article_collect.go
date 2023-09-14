package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
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
	var num = -1
	if err != nil {
		// not found
		global.DB.Create(&models.UserArticleCollectionModels{
			UserID:    claim.UserID,
			ArticleID: req.ID,
		})
		// collect
		num = 1
	}
	// suspend collect
	global.DB.Delete(&collection)

	// update collect count
	es_service.ArticleUpdate(req.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		res.OkWithMessage("Collect article successfully", c)
	} else {
		res.OkWithMessage("Failed to collect article", c)
	}
}
