package setting_api

import (
	"blog_server/global"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingView(c *gin.Context) {
	var uri SiteUri
	err := c.ShouldBindUri(&uri)

	if err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}

	switch uri.Name {
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "aws":
		res.OkWithData(global.Config.AWS, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("No such information to get.", c)
	}
}
