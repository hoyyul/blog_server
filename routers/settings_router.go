package routers

import "blog_server/api"

func (rg RouterGroup) SettingsRouter() {
	settingsApp := api.ApiGroupApp.SettingsApi
	rg.GET("settings/:name", settingsApp.SettingsGetInfoView)
	rg.PUT("settings/:name", settingsApp.SettingsUpdateInfoView)
}
