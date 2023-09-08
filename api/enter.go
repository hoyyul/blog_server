package api

import (
	"blog_server/api/advertise_api"
	"blog_server/api/image_api"
	"blog_server/api/menu_api"
	"blog_server/api/message_api"
	"blog_server/api/setting_api"
	"blog_server/api/tag_api"
	"blog_server/api/user_api"
)

type ApiGroup struct {
	SettingApi   setting_api.SettingApi
	ImageApi     image_api.ImageApi
	AdvertiseApi advertise_api.AdvertiseApi
	MenuApi      menu_api.MenuApi
	UserApi      user_api.UserApi
	TagApi       tag_api.TagApi
	MessageApi   message_api.MessageApi
}

var ApiGroupApp = new(ApiGroup)
