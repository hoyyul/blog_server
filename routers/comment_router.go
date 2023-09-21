package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	rg.POST("comment", middleware.CheckAuthToken(), app.CommentCreateView)
	//rg.GET("comments", app.CommentListView)
}
