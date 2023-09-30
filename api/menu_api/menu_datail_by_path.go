package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type MenuDetailRequest struct {
	Path string `json:"path" form:"path"`
}

func (MenuApi) MenuDetailByPathView(c *gin.Context) {
	var req MenuDetailRequest
	err := c.ShouldBindQuery(&req) // api/menus/detail?path=/
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, "path = ?", req.Path).Error
	if err != nil {
		res.FailWithMessage("Menu doesn't exit", c)
		return
	}

	var menuBanners []models.MenuBannerModel
	global.DB.Preload(clause.Associations).Order("sort desc").Find(&menuBanners, "menu_id = ?", menuModel.ID)
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.Banner.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
}
