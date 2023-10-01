package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ImageRouter() {
	app := api.ApiGroupApp.ImageApi
	rg.POST("images", middleware.CheckAdminToken(), app.ImageUploadListView)
	rg.POST("image", middleware.CheckAdminToken(), app.ImageUploadView)
	rg.GET("images", app.ImageListView)
	rg.GET("image_names", app.ImageNameListView)
	rg.PUT("images", middleware.CheckAdminToken(), app.ImageUpdateName)
	rg.DELETE("images", middleware.CheckAdminToken(), app.ImageRemoveListView)
}
