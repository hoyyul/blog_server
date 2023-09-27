package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ImageRouter() {
	app := api.ApiGroupApp.ImageApi
	rg.POST("images", middleware.CheckAdminToken(), app.ImageCreateListView)
	rg.POST("image", middleware.CheckAdminToken(), app.ImageUploadView)
	rg.GET("images", app.ImageReadListView)
	rg.GET("image_names", app.ImageReadNameListView)
	rg.PUT("images", middleware.CheckAdminToken(), app.ImageUpdateName)
	rg.DELETE("images", middleware.CheckAdminToken(), app.ImageDeleteListView)
}
