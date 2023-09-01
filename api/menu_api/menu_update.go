package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var req MenuRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	id := c.Param("id")

	// clean original banners
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("Menu doesn't exits", c)
		return
	}

	global.DB.Model(&menuModel).Association("Banners").Clear() // clean banners in banner-menu table

	// Add new banners to banner-menu table
	if len(req.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, sort := range req.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage("Failed to add new banners", c)
			return
		}
	}

	// update the other fielddata
	maps := structs.Map(&req)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update menu", c)
		return
	}

	res.OkWithMessage("Update menu successfully", c)

}
