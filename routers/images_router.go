package routers

import "blog_server/api"

func (rg RouterGroup) ImagesRouter() {
	imagesApp := api.ApiGroupApp.ImagesApi
	rg.POST("images", imagesApp.ImagesUploadView)
	rg.GET("images", imagesApp.ImagesGetListView)
	rg.GET("image_names", imagesApp.ImageGetNameListView)
	rg.DELETE("images", imagesApp.ImagesRemoveListView)
	rg.PUT("images", imagesApp.ImagesUpdateName)
}
