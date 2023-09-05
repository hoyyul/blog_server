package user_api

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service"
	"blog_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	_claim, _ := c.Get("claim")
	claim := _claim.(*jwts.CustomClaim)

	token := c.Request.Header.Get("token")

	err := service.ServiceApp.UserService.Logout(claim, token)

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage("Failed to logout", c)
		return
	}

	res.OkWithMessage("Logout successfully", c)

}
