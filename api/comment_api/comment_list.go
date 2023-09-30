package comment_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/redis_service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"id" uri:"id" json:"id"`
}

func (CommentApi) CommentListView(c *gin.Context) {
	var req CommentListRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	rootCommentList := FindArticleCommentList(req.ArticleID)
	res.OkWithData(filter.Select("c", rootCommentList), c)
}

func FindArticleCommentList(articleID string) (rootCommentList []*models.CommentModel) {
	// find all parent comments
	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)

	// find all subcomments of each parent comment
	diggInfo := redis_service.NewCommentDigg().GetInfo()
	for _, model := range rootCommentList {
		var subCommentList, newSubCommentList []models.CommentModel
		Recursion(*model, &subCommentList)
		for _, commentModel := range subCommentList {
			digg := diggInfo[fmt.Sprintf("%d", commentModel.ID)]
			commentModel.DiggCount = commentModel.DiggCount + digg
			newSubCommentList = append(newSubCommentList, commentModel)
		}
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		model.DiggCount = model.DiggCount + modelDigg
		model.SubComments = newSubCommentList
	}
	return
}

func FindSubCommentList(model models.CommentModel) (subCommentList []models.CommentModel) {
	Recursion(model, &subCommentList)
	return subCommentList
}

func Recursion(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, subcomment := range model.SubComments {
		*subCommentList = append(*subCommentList, subcomment)
		Recursion(subcomment, subCommentList)
	}
}
