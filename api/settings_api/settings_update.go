package settings_api

import (
	"blog_server/config"
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsUpdateInfoView(c *gin.Context) {
	var cs config.SiteInfo
	err := c.ShouldBindJSON(&cs)

	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	global.Config.SiteInfo = cs
	err = initialization.SettingYaml()

	if err != nil {
		global.Logger.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	global.Logger.Info("Update configuration file successfully!")
	res.OkWithSuccess(c)
}
