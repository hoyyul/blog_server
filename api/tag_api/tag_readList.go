package tag_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common_service"

	"github.com/gin-gonic/gin"
)

func (TagApi) TagReadListView(c *gin.Context) {
	var req models.PageInfo
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	list, count, _ := common_service.FetchPaginatedData[models.TagModel](models.TagModel{}, common_service.Option{
		PageInfo: req,
	})

	res.OkWithList(list, count, c)
}
