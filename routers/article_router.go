package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	rg.POST("article", middleware.CheckAuthToken(), app.ArticleCreateView)
	rg.GET("article", app.ArticleReadListView)
	rg.GET("article/:id", app.ArticleReadDetailView)
	rg.GET("article/detail", app.ArticleReadDetailByTitleView)
}
