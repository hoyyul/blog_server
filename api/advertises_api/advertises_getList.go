package advertises_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common_service"
	"strings"

	"github.com/gin-gonic/gin"
)

func (AdvertiseApi) AdvertisesGetListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)

	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// if request came from admin, return all; else return isShow = true
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}

	// get paginated advertisement list
	advertisesList, count, err := common_service.FetchPaginatedData[models.AdvertiseModel](models.AdvertiseModel{IsShow: isShow}, common_service.Option{
		PageInfo: page,
		Debug:    true,
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithList(advertisesList, count, c)
}
