package comment_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common_service"
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CommentListResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	ArticleTitle    string    `json:"article_title"`
	ArticleBanner   string    `json:"article_banner"`
	ParentCommentID *uint     `json:"parent_comment_id"`
	Content         string    `json:"content"`
	DiggCount       int       `json:"digg_count"`
	CommentCount    int       `json:"comment_count"`
	UserNickName    string    `json:"user_nick_name"`
}

func (CommentApi) CommentListAllView(c *gin.Context) {
	var req models.PageInfo

	c.ShouldBindQuery(&req)

	list, count, _ := common_service.FetchPaginatedData[models.CommentModel](models.CommentModel{}, common_service.Option{
		PageInfo: req,
		Preload:  []string{"User"},
	})

	var commentList = make([]CommentListResponse, 0)

	var collMap = map[string]models.ArticleModel{}
	var articleIDList []interface{}

	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
	}
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		collMap[hit.Id] = article
	}

	for _, model := range list {
		commentList = append(commentList, CommentListResponse{
			ID:              model.ID,
			CreatedAt:       model.CreatedAt,
			ParentCommentID: model.ParentCommentID,
			Content:         model.Content,
			DiggCount:       model.DiggCount,
			CommentCount:    model.CommentCount,
			UserNickName:    model.User.NickName,
			ArticleTitle:    collMap[model.ArticleID].Title,
			ArticleBanner:   collMap[model.ArticleID].BannerUrl,
		})
	}

	res.OkWithList(commentList, count, c)

}
