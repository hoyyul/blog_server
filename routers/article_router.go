package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	rg.POST("article", middleware.CheckAuthToken(), app.ArticleCreateView)
	rg.POST("article/collect", middleware.CheckAuthToken(), app.ArticleCollCreateView)
	rg.GET("article", app.ArticleReadListView)
	rg.GET("article/:id", app.ArticleReadDetailView)
	rg.GET("article/detail", app.ArticleReadDetailByTitleView)
	rg.GET("article/calendar", app.ArticleReadCalendarCountView)
	rg.PUT("article", app.ArticleUpdateView)
	rg.DELETE("article", app.ArticleRemoveView)
	rg.DELETE("article/collect", middleware.CheckAuthToken(), app.ArticleDeleteCollectView)
}
