package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	rg.POST("article", middleware.CheckAuthToken(), app.ArticleCreateView)
	rg.GET("article", app.ArticleReadListView)
}
