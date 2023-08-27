package routers

import "blog_server/api"

func (rg RouterGroup) SettingsRouter() {
	settingsApp := api.ApiGroupApp.SettingsApi
	rg.GET("settings", settingsApp.SettingsInfoView)
	rg.PUT("settings", settingsApp.SettingsUpdateInfoView)
}
