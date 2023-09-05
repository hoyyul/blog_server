package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"Enter old password"`
	Pwd    string `json:"pwd" binding:"required" msg:"Enter a new password"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	var req UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithValidation(err, &req, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claim.UserID).Error
	if err != nil {
		res.FailWithMessage("User doesn't exist", c)
		return
	}

	if !pwd.CheckPwd(user.Password, req.OldPwd) {
		res.FailWithMessage("Password incorrect", c)
		return
	}
	hashPwd := pwd.HashPwd(req.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to change password", c)
		return
	}
	res.OkWithMessage("Change password successfullt", c)
}
