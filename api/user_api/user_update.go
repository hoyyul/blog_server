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
	var req UserRole
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}

	var user models.UserModel
	err := global.DB.Take(&user, req.UserID).Error
	if err != nil {
		res.FailWithMessage("User doesn't exits", c)
		return
	}

	// update user information
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      req.Role,
		"nick_name": req.NickName,
	}).Error

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update user information", c)
		return
	}
	res.OkWithMessage("Update user information successfully", c)
}
