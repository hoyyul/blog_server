package routers

import "blog_server/api"

func (rg RouterGroup) AdvertiseRouter() {
	app := api.ApiGroupApp.AdvertiseApi
	rg.POST("advertisement", app.AdvertiseCreateView)
	rg.GET("advertisement", app.AdvertiseReadListView)
	rg.PUT("advertisement/:id", app.AdvertiseUpdateView)
	rg.DELETE("advertisement", app.AdvertiseDeletView)
}
