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
	Title         string      `json:"title" binding:"required" msg:"Enter a menu title" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"Enter a menu path" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"` // time interval to swicth images
	BannerTime    int         `json:"banner_time" structs:"banner_time"`
	Sort          int         `json:"sort" structs:"sort"`         // menu sort
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"` // image sort
}

func (MenuApi) MenuUploadView(c *gin.Context) {
	var req MenuRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	// check if menu already exists
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", req.Title, req.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("Menu already exists", c)
		return
	}

	// save data to menu table
	menuModel := models.MenuModel{
		Title:        req.Title,
		Path:         req.Path,
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
