package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	rg.POST("tags", middleware.CheckAdminToken(), app.TagUploadView)
	rg.GET("tags", app.TagListView)
	rg.GET("tag_names", app.TagNameListView)
	rg.PUT("tags/:id", middleware.CheckAdminToken(), app.TagUpdateView)
	rg.DELETE("tags", middleware.CheckAdminToken(), app.TagRemoveListView)
}
