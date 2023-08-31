package advertises_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

// AdvertisesRemoveView Remove Advertisements
// @Tags Advertisement Management
// @Summary Remove Advertisements
// @Description Remove Advertisements
// @Param data body models.RemoveRequest    true  "Advertisement idList"
// @Router /api/advertisement [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertiseApi) AdvertisesRemoveView(c *gin.Context) {
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
