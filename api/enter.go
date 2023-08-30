package api

import (
	"blog_server/api/advertises_api"
	"blog_server/api/images_api"
	"blog_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi   settings_api.SettingsApi
	ImagesApi     images_api.ImagesApi
	AdvertisesApi advertises_api.AdvertiseApi
}

var ApiGroupApp = new(ApiGroup)
