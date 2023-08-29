package images_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"

	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImagesGetListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)

	if err != nil {
		res.FailWithCode(res.ParameterError, c)
	}

	// get paginated image list
	imageList, count, err := common.FetchPaginatedData[models.BannerModel](common.Option{
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
