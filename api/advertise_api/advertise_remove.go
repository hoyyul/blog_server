package advertise_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (AdvertiseApi) AdvertiseRemoveView(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	var advertiseList []models.AdvertiseModel
	count := global.DB.Find(&advertiseList, req.IdList).RowsAffected
	if count == 0 {
		res.FailWithMessage("Advertisement doesn't exist", c)
		return
	}
	global.DB.Delete(&advertiseList)
	res.OkWithMessage(fmt.Sprintf("Remove %d advertisements", count), c)

}
