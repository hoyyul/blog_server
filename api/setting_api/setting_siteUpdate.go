package setting_api

import (
	"blog_server/config"
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	global.Config.SiteInfo = info
	initialization.SettingYaml()
	res.OkWithMessage("Update site info successfully", c)
}
