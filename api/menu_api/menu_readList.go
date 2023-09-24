package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuReadListView(c *gin.Context) {
	// get menu
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList) // scan database type to string type, "myblog\nforfun"->["myblog", "for fun"]
	// get menu-banner model
	var menuBanners []models.MenuBannerModel
	global.DB.Preload(clause.Associations).Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus = make([]MenuResponse, 0)
	for _, model := range menuList {
		var banners = make([]Banner, 0)
		for _, record := range menuBanners {
			if model.ID == record.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   record.BannerID,
				Path: record.Banner.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
}
