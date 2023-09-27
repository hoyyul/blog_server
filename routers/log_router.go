package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	rg.GET("logs", app.LogReadListView)
	rg.DELETE("logs", middleware.CheckAdminToken(), app.LogDeleteListView)
}
