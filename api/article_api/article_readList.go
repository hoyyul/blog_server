package article_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/es_service"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ArticleApi) ArticleReadListView(c *gin.Context) {
	var req models.PageInfo
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	list, count, err := es_service.GetList(req.Key, req.Page, req.Limit)
	if err != nil {
		global.Logger.Error(err)
		res.OkWithMessage("Failed to get list", c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
