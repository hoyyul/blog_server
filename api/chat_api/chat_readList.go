package chat_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common_service"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ChatApi) ChatReadListView(c *gin.Context) {
	var req models.PageInfo
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	req.Sort = "created_at desc"
	list, count, _ := common_service.FetchPaginatedData[models.ChatModel](models.ChatModel{IsGroup: true}, common_service.Option{
		PageInfo: req,
	})

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	res.OkWithList(data, count, c)
}
