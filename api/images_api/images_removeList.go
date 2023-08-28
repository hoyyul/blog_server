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
	count := global.DB.Find(&imageList, req.IdList).RowsAffected // search
	if count == 0 {
		res.FailWithMessage("No image can be removed", c)
		return
	}

	global.DB.Delete(imageList) // delete
	res.OkWithMessage(fmt.Sprintf("Remove %d image", count), c)
}
