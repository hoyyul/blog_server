package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	rg.POST("messages", app.MessageCreateView)
	rg.GET("messages_all", app.MessageReadListView)
	rg.GET("messages", middleware.CheckAuthToken(), app.MessageReadHistoryView)
	rg.GET("messages_record", middleware.CheckAuthToken(), app.MessageReadRecordView)
}
