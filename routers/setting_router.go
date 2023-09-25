package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) SettingRouter() {
	app := api.ApiGroupApp.SettingApi
	rg.GET("setting/:name", app.SettingReadView)
	rg.PUT("setting/:name", middleware.CheckAdminToken(), app.SettingUpdateView)
}
