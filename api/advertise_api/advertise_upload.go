package advertise_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (AdvertiseApi) AdvertiseUploadView(c *gin.Context) {
	var req AdvertiseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, req, c)
		return
	}

	var advertisement models.AdvertiseModel
	err = global.DB.Take(&advertisement, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMessage("Advertisement already exists", c)
		return
	}

	err = global.DB.Create(&models.AdvertiseModel{
		Title:  req.Title,
		Href:   req.Href,
		Images: req.Images,
		IsShow: req.IsShow,
	}).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to add advertisement", c)
		return
	}

	res.OkWithMessage("Add advertisement successfully", c)
}
