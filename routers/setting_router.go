package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) SettingRouter() {
	app := api.ApiGroupApp.SettingApi
	rg.GET("settings/site", app.SettingSiteInfoView)
	rg.GET("settings/:name", middleware.CheckAdminToken(), app.SettingView)
	rg.PUT("settings/site", middleware.CheckAdminToken(), app.SettingSiteUpdateView)
	rg.PUT("settings/:name", middleware.CheckAdminToken(), app.SettingUpdateView)
}
