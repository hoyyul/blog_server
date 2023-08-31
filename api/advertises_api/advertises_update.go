package advertises_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertisesUpdateView Update Advertisement
// @Tags Advertisement Management
// @Summary Update Advertisement
// @Description Update Advertisement
// @Param data body AdvertiseRequest   true  "some parameters"
// @Router /api/advertisement/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertiseApi) AdvertisesUpdateView(c *gin.Context) {
	var req AdvertiseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, req, c)
		return
	}

	var advertisement models.AdvertiseModel
	err = global.DB.Take(&advertisement, c.Param("id")).Error
	if err != nil {
		res.FailWithMessage("Advertisement doesn't exist", c)
		return
	}

	// convert struct to map
	maps := structs.Map(&req)
	err = global.DB.Model(&advertisement).Updates(maps).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update advertisement", c)
		return
	}

	res.OkWithMessage("Update advertisement successfully", c)
}
