package api

import (
	"blog_server/api/advertise_api"
	"blog_server/api/image_api"
	"blog_server/api/menu_api"
	"blog_server/api/setting_api"
)

type ApiGroup struct {
	SettingApi   setting_api.SettingApi
	ImageApi     image_api.ImageApi
	AdvertiseApi advertise_api.AdvertiseApi
	MenuApi      menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
