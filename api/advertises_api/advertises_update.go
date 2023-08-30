package advertises_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

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

	err = global.DB.Model(&advertisement).Updates(map[string]any{
		"title":   req.Title,
		"href":    req.Href,
		"images":  req.Images,
		"is_show": req.IsShow,
	}).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update advertisement", c)
		return
	}

	res.OkWithMessage("Update advertisement successfully", c)
}
