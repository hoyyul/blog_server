package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"Enter menu title"`
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"Enter English menu title"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"`
	BannerTime    int         `json:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"Enter menu sort"`
	ImageSortList []ImageSort `json:"image_sort_list"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var req MenuRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	// save data to menu table
	menuModel := models.MenuModel{
		MenuTitle:    req.MenuTitle,
		MenuTitleEn:  req.MenuTitleEn,
		Slogan:       req.Slogan,
		Abstract:     req.Abstract,
		AbstractTime: req.AbstractTime,
		BannerTime:   req.BannerTime,
		Sort:         req.Sort,
	}

	err = global.DB.Create(&menuModel).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to add menu", c)
		return
	}
	if len(req.ImageSortList) == 0 {
		res.OkWithMessage("Add menu successfully", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range req.ImageSortList {
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// save data to menu-banner table
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update menu-banner table", c)
		return
	}
	res.OkWithMessage("Add menu successfully", c)
}
