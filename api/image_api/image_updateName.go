package image_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

/*
Only change image name in database, without changing path in local
*/
func (ImageApi) ImageUpdateName(c *gin.Context) {
	var req ImageUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	var image models.BannerModel
	// check if image exist
	err = global.DB.Take(&image, req.ID).Error
	if err != nil {
		res.FailWithMessage("Image doesn't exist", c)
	}

	// update image name
	err = global.DB.Model(&image).Update("name", req.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage("Update image name successfully!", c)
}
