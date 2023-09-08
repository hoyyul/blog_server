package tag_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (TagApi) TagDeleteListView(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, req.IdList).RowsAffected
	if count == 0 {
		res.FailWithMessage("Tag doesn't exist", c)
		return
	}

	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("Deleted %d tags", count), c)

}
