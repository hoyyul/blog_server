package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	rg.POST("comments", middleware.CheckAuthToken(), app.CommentCreateView)
	rg.GET("comments/digg/:id", app.CommentDiggView)
	rg.GET("comments", app.CommentListAllView)
	rg.GET("comments/:id", app.CommentReadListView) // get comments of an article
	rg.DELETE("comments/:id", middleware.CheckAdminToken(), app.CommentDeleteView)

}
