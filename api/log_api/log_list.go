package log_api

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugins/log_stash"
	"blog_server/service/common_service"

	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `form:"level"`
}

func (LogApi) LogListView(c *gin.Context) {
	var req LogRequest
	c.ShouldBindQuery(&req)
	list, count, _ := common_service.FetchPaginatedData[log_stash.LogStashModel](log_stash.LogStashModel{Level: req.Level}, common_service.Option{
		PageInfo: req.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return
}
