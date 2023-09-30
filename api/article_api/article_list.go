package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"
	"blog_server/service/redis_service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag    string `json:"tag" form:"tag"`
	IsUser bool   `json:"is_user" form:"is_user"` // if user collected
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var req ArticleSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	boolSearch := elastic.NewBoolQuery()

	if req.IsUser {
		token := c.GetHeader("token")
		claims, err := jwts.ParseToken(token)
		if err == nil && !redis_service.CheckLogout(token) {
			boolSearch.Must(elastic.NewTermsQuery("user_id", claims.UserID))
		}
	}

	// paginate search by title + tag
	list, count, err := es_service.GetList(es_service.Option{
		PageInfo: req.PageInfo,
		Fields:   []string{"title", "content", "category"},
		Tag:      req.Tag,
		Query:    boolSearch,
	})
	if err != nil {
		global.Logger.Error(err)
		res.OkWithMessage("Failed to get list", c)
		return
	}

	// list can't be {}
	data := filter.Omit("list", list) // ignore to present field with "omit(list)"
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" { // can't find such a article
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	res.OkWithList(data, int64(count), c)
}
