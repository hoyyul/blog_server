package comment_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/service/redis_service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"Enter article id"`
	Content         string `json:"content" binding:"required" msg:"Enter article comment"`
	ParentCommentID *uint  `json:"parent_comment_id"` // parent id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var req CommentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	// if article exists
	_, err = es_service.GetDetail(req.ArticleID)
	if err != nil {
		res.FailWithMessage("Article doesn't exist", c)
		return
	}

	// if subcomment
	if req.ParentCommentID != nil {
		var parentComment models.CommentModel

		// find parent
		err = global.DB.Take(&parentComment, req.ParentCommentID).Error
		if err != nil {
			res.FailWithMessage("Parent comment doesn't exist", c)
			return
		}
		if parentComment.ArticleID != req.ArticleID {
			res.FailWithMessage("Parent ID incorrect", c)
			return
		}
		// add one comment count to parent ID
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}

	// add comment to database
	global.DB.Create(&models.CommentModel{
		ParentCommentID: req.ParentCommentID,
		Content:         req.Content,
		ArticleID:       req.ArticleID,
		UserID:          claim.UserID,
	})

	// add one total comment count to redis
	redis_service.NewCommentCount().Set(req.ArticleID)
	res.OkWithMessage("Publish comment successfully", c)
}
