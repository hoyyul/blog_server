package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	rg.POST("tags", middleware.CheckAdminToken(), app.TagCreateView)
	rg.GET("tags", app.TagReadListView)
	rg.GET("tag_names", app.TagNameListView)
	rg.PUT("tags/:id", middleware.CheckAdminToken(), app.TagUpdateView)
	rg.DELETE("tags", middleware.CheckAdminToken(), app.TagDeleteListView)
}
