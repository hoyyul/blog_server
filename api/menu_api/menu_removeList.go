package menu_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveListView(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, req.IdList).RowsAffected
	if count == 0 {
		res.FailWithMessage("Menu doesn't exist", c)
		return
	}
	fmt.Println("idlist", req.IdList)
	// transaction to delete data in menu-banner table and menu table
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to delete menu", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("Remove %d menus", count), c)

}
