package routers

import "blog_server/api"

func (router RouterGroup) AdvertisesRouter() {
	app := api.ApiGroupApp.AdvertisesApi
	router.POST("advertisement", app.AdvertisementUploadView)
	router.GET("advertisement", app.AdvertisesGetListView)
	router.PUT("advertisement/:id", app.AdvertisesUpdateView)
}
