package routers

import "blog_server/api"

func (router RouterGroup) AdvertiseRouter() {
	app := api.ApiGroupApp.AdvertiseApi
	router.POST("advertisement", app.AdvertiseCreateView)
	router.GET("advertisement", app.AdvertiseReadListView)
	router.PUT("advertisement/:id", app.AdvertiseUpdateView)
	router.DELETE("advertisement", app.AdvertiseDeletView)
}
