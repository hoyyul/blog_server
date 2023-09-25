package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	rg.POST("comment", middleware.CheckAuthToken(), app.CommentCreateView)
	rg.POST("comment/:id", app.CommentDiggView)
	rg.GET("comment", app.CommentReadListView)
	rg.DELETE("comment/:id", middleware.CheckAdminToken(), app.CommentDeleteView)
}
