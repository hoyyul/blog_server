package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticleApi) ArticleReadListView(c *gin.Context) {
	var req ArticleSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// paginate search by title + tag
	list, count, err := es_service.GetList(es_service.Option{
		PageInfo: req.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      req.Tag,
	})
	if err != nil {
		global.Logger.Error(err)
		res.OkWithMessage("Failed to get list", c)
		return
	}

	// list can't be {}
	data := filter.Omit("list", list) // ignore field with "omit(list)"
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	res.OkWithList(data, int64(count), c)
}
