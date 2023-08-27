package settings_api

import (
	"blog_server/config"
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsUpdateInfoView(c *gin.Context) {
	var uri SiteUri
	err := c.ShouldBindUri(&uri) // load request uri into struct

	if err != nil {
		global.Logger.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}

	configMap := map[string]interface{}{
		"site":  &config.SiteInfo{},
		"email": &config.Email{},
		"qq":    &config.QQ{},
		"qiniu": &config.QiNiu{},
		"jwt":   &config.Jwt{},
	}

	if info, ok := configMap[uri.Name]; ok {
		err = c.ShouldBindJSON(info)

		if err != nil {
			global.Logger.Error(err)
			res.FailWithCode(res.ParameterError, c)
			return
		}

		// judge which entry to be updated
		switch uri.Name {
		case "site":
			global.Config.SiteInfo = *(info.(*config.SiteInfo)) //reflection
		case "email":
			global.Config.Email = *(info.(*config.Email))
		case "qq":
			global.Config.QQ = *(info.(*config.QQ))
		case "qiniu":
			global.Config.QiNiu = *(info.(*config.QiNiu))
		case "jwt":
			global.Config.Jwt = *(info.(*config.Jwt))
		}

		err = initialization.SettingYaml()

		if err != nil {
			global.Logger.Error(err)
			res.FailWithMessage(err.Error(), c)
			return
		}

		global.Logger.Info("Update configuration file successfully!")
		res.OkWithSuccess(c)

	} else {
		res.FailWithMessage("No such information to update.", c)
	}
}
