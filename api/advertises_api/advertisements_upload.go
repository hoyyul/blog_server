package advertises_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

type AdvertiseRequest struct {
	Title  string `json:"title" binding:"required" msg:"Enter a title"`
	Href   string `json:"href" binding:"required,url" msg:"Enter a valid url"`
	Images string `json:"images" binding:"required,url" msg:"Enter a valid image path"`
	IsShow bool   `json:"is_show" binding:"required" msg:"Select show or not"`
}

func (AdvertiseApi) AdvertisementUploadView(c *gin.Context) {
	var req AdvertiseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, req, c)
		return
	}

	var advert models.AdvertiseModel
	err = global.DB.Take(&advert, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMessage("This advertisements already exists", c)
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
