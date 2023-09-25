package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	rg.POST("tag", middleware.CheckAdminToken(), app.TagCreateView)
	rg.GET("tag", app.TagReadListView)
	rg.PUT("tag/:id", middleware.CheckAdminToken(), app.TagUpdateView)
	rg.DELETE("tag", middleware.CheckAdminToken(), app.TagDeleteListView)
}
