package comment_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/redis_service"
	"blog_server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var req CommentIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, req.ID).Error
	if err != nil {
		res.FailWithMessage("Comment doesn't exist", c)
		return
	}

	// count = sum(subcomment) + 1(self)
	subCommentList := FindSubCommentList(commentModel)
	count := len(subCommentList) + 1
	redis_service.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	// subcomment?
	if commentModel.ParentCommentID != nil {
		// find parent comment and deduct comment count
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// delete subcomment and self
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}

	// reverse
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OkWithMessage(fmt.Sprintf("Deleted %d comments", len(deleteCommentIDList)), c)
}
