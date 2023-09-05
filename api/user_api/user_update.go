package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"Permission Parameter Error"`
	NickName string     `json:"nick_name"`
	UserID   uint       `json:"user_id" binding:"required" msg:"Enter userid"`
}

func (UserApi) UserUpdateView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidation(err, &cr, c)
		return
	}

	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("User doesn't exits", c)
		return
	}

	// update user information
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update user information", c)
		return
	}
	res.OkWithMessage("Update user information successfully", c)
}
