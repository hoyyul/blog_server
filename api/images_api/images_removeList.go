package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImagesRemoveListView(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	var imageList []models.BannerModel
	// check if data exists before deletion
	count := global.DB.Find(&imageList, req.IdList).RowsAffected
	if count == 0 {
		res.FailWithMessage("No image can be removed", c)
		return
	}

	// delete image in database, and use hook function to delect image in local
	global.DB.Delete(imageList)
	res.OkWithMessage(fmt.Sprintf("Remove %d image", count), c)
}
