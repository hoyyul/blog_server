package log_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugins/log_stash"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (LogApi) LogRemoveListView(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	var list []log_stash.LogStashModel
	count := global.DB.Find(&list, req.IdList).RowsAffected
	if count == 0 {
		res.FailWithMessage("Log doesn't exist", c)
		return
	}
	global.DB.Delete(&list)
	res.OkWithMessage(fmt.Sprintf("Deleted %d logs", count), c)

}
