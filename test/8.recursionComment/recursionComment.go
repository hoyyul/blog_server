package main

import (
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/models"
	"fmt"
)

func main() {
	initialization.InitConf()
	global.Logger = initialization.InitLogger()
	global.DB = initialization.InitGorm()
	FindArticleCommentList("y4A1sooBTckTgKfb4Yr8")
}

func FindArticleCommentList(articleID string) {
	var RootCommentList []*models.CommentModel

	// find all parent comments
	global.DB.Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)

	// find all subcomments of each parent comment
	for _, model := range RootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
		fmt.Println(subCommentList)
	}
}

func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&model)
	for _, subcomment := range model.SubComments {
		*subCommentList = append(*subCommentList, subcomment)
		FindSubComment(subcomment, subCommentList)
	}
}
