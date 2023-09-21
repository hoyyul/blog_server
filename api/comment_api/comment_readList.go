package comment_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

func (CommentApi) CommentReadListView(c *gin.Context) {
	var req CommentListRequest
	err := c.ShouldBindQuery(&req)
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
	for _, model := range rootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
		//fmt.Println(subCommentList)
	}
	return
}

func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, subcomment := range model.SubComments {
		*subCommentList = append(*subCommentList, subcomment)
		FindSubComment(subcomment, subCommentList)
	}
}
