package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	rg.POST("messages", app.MessageUploadView)
	rg.GET("messages_all", app.MessageListAllView)
	rg.GET("messages", middleware.CheckAuthToken(), app.MessageListView)
	rg.GET("messages_record", middleware.CheckAuthToken(), app.MessageRecordView)
}
