package setting_api

import (
	"blog_server/global"
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingSiteInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
