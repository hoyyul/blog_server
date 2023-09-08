package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	rg.POST("message", app.MessageCreateView)
	rg.GET("message_list", app.MessageReadListView)
	rg.GET("message", middleware.CheckAuthToken(), app.MessageReadHistoryView)
}
