package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) AdvertiseRouter() {
	app := api.ApiGroupApp.AdvertiseApi
	rg.POST("adverts", middleware.CheckAdminToken(), app.AdvertiseCreateView)
	rg.GET("adverts", app.AdvertiseReadListView)
	rg.PUT("adverts/:id", middleware.CheckAdminToken(), app.AdvertiseUpdateView)
	rg.DELETE("adverts", middleware.CheckAdminToken(), app.AdvertiseDeletView)
}
