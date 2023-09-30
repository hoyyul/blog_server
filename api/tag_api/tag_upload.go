package tag_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"Enter a title" structs:"title"`
}

func (TagApi) TagUploadView(c *gin.Context) {
	var req TagRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMessage("Tag already exists", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: req.Title,
	}).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to add tag", c)
		return
	}

	res.OkWithMessage("Add tag successfullt", c)
}
