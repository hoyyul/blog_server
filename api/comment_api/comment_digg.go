package comment_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/redis_service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"` // comment id
}

func (CommentApi) CommentDiggView(c *gin.Context) {
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

	redis_service.NewCommentDigg().Set(fmt.Sprintf("%d", req.ID))

	res.OkWithMessage("Digged comment successfully", c)
}
