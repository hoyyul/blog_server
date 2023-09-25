package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ImageRouter() {
	app := api.ApiGroupApp.ImageApi
	rg.POST("image", middleware.CheckAdminToken(), app.ImageCreateListView)
	rg.GET("image", app.ImageReadListView)
	rg.GET("image_name", app.ImageReadNameListView)
	rg.PUT("image", middleware.CheckAdminToken(), app.ImageUpdateName)
	rg.DELETE("image", middleware.CheckAdminToken(), app.ImageDeleteListView)
}
