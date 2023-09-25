package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) AdvertiseRouter() {
	app := api.ApiGroupApp.AdvertiseApi
	rg.POST("advertisement", middleware.CheckAdminToken(), app.AdvertiseCreateView)
	rg.GET("advertisement", app.AdvertiseReadListView)
	rg.PUT("advertisement/:id", middleware.CheckAdminToken(), app.AdvertiseUpdateView)
	rg.DELETE("advertisement", middleware.CheckAdminToken(), app.AdvertiseDeletView)
}
