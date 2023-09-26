package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
}

func (UserApi) UserUpdateNickName(c *gin.Context) {
	var req UserUpdateNicknameRequest
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var newMaps = map[string]interface{}{}
	maps := structs.Map(req)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}
	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, claim.UserID).Error
	if err != nil {
		res.FailWithMessage("User doesn't exist", c)
		return
	}
	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to update user info", c)
		return
	}
	res.OkWithMessage("Update user info successfully", c)

}
