package api

import (
	"blog_server/api/advertise_api"
	"blog_server/api/article_api"
	"blog_server/api/chat_api"
	"blog_server/api/comment_api"
	"blog_server/api/image_api"
	"blog_server/api/log_api"
	"blog_server/api/menu_api"
	"blog_server/api/message_api"
	"blog_server/api/news_api"
	"blog_server/api/setting_api"
	"blog_server/api/statistic_api"
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
	ArticleApi   article_api.ArticleApi
	//DiggApi      digg_api.DiggApi
	CommentApi   comment_api.CommentApi
	NewsApi      news_api.NewsApi
	ChatApi      chat_api.ChatApi
	LogApi       log_api.LogApi
	StatisticApi statistic_api.StatisticApi
}

var ApiGroupApp = new(ApiGroup)
