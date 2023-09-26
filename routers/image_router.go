package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ImageRouter() {
	app := api.ApiGroupApp.ImageApi
	rg.POST("images", middleware.CheckAdminToken(), app.ImageCreateListView)
	rg.POST("image", middleware.CheckAdminToken(), app.ImageUploadView)
	rg.GET("image", app.ImageReadListView)
	rg.GET("image_name", app.ImageReadNameListView)
	rg.PUT("image", middleware.CheckAdminToken(), app.ImageUpdateName)
	rg.DELETE("image", middleware.CheckAdminToken(), app.ImageDeleteListView)
}
