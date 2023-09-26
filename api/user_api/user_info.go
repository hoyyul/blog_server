package user_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (UserApi) UserInfoView(c *gin.Context) {

	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	var userInfo models.UserModel
	err := global.DB.Take(&userInfo, claim.UserID).Error
	if err != nil {
		res.FailWithMessage("User doesn't exist", c)
		return
	}
	res.OkWithData(filter.Select("info", userInfo), c)

}
