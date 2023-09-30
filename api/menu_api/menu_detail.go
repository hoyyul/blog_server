package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	// search menu table
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("menu doesn't exist", c)
		return
	}
	// search menu-banner table
	var menuBanners []models.MenuBannerModel
	global.DB.Preload(clause.Associations).Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var banners = make([]Banner, 0)
	for _, record := range menuBanners {
		if menuModel.ID != record.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   record.BannerID,
			Path: record.Banner.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
}
