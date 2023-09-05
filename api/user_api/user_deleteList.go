package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserDeleteListView(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, req.IdList).RowsAffected
	if count == 0 {
		res.FailWithMessage("User doesn't exist", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO: delect user, collections, publish...
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Logger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to remove user", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("Removed %d users", count), c)

}
