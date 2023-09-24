package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	rg.GET("log", app.LogReadListView)
	rg.DELETE("log", middleware.CheckAdminToken(), app.LogDeleteListView)
}
