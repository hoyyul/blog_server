package tag_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdateView(c *gin.Context) {

	id := c.Param("id")
	var req TagRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("Tag doesn't exist", c)
		return
	}

	maps := structs.Map(&req)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update tag", c)
		return
	}

	res.OkWithMessage("Updated tag successfully", c)
}
