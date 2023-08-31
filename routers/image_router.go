package routers

import "blog_server/api"

func (rg RouterGroup) ImageRouter() {
	app := api.ApiGroupApp.ImageApi
	rg.POST("image", app.ImageCreateListView)
	rg.GET("image", app.ImageReadListView)
	rg.GET("image_name", app.ImageReadNameListView)
	rg.PUT("image", app.ImageUpdateName)
	rg.DELETE("image", app.ImageDeleteListView)
}
