package image_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common_service"

	"github.com/gin-gonic/gin"
)

func (ImageApi) ImageListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)

	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// get paginated image list
	imageList, count, err := common_service.FetchPaginatedData[models.BannerModel](models.BannerModel{}, common_service.Option{
		PageInfo: page,
		Debug:    true,
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithList(imageList, count, c)
}
