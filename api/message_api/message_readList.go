package message_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common_service"

	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageReadListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	list, count, _ := common_service.FetchPaginatedData[models.MessageModel](models.MessageModel{}, common_service.Option{
		PageInfo: cr,
	})

	res.OkWithList(list, count, c)
}
