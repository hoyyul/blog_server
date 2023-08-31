package routers

import "blog_server/api"

func (rg RouterGroup) SettingRouter() {
	app := api.ApiGroupApp.SettingApi
	rg.GET("setting/:name", app.SettingReadView)
	rg.PUT("setting/:name", app.SettingUpdateView)
}
