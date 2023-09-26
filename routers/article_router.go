package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	rg.POST("articles", middleware.CheckAdminToken(), app.ArticleCreateView)
	rg.POST("articles/collects", middleware.CheckAuthToken(), app.ArticleCollCreateView)
	rg.GET("categorys", app.ArticleCategoryListView)
	rg.GET("articles/content/:id", app.ArticleContentByIDView)
	rg.GET("articles", app.ArticleReadListView)
	rg.GET("articles/:id", app.ArticleReadDetailView)
	rg.GET("articles/detail", app.ArticleReadDetailByTitleView)
	rg.GET("articles/calendar", app.ArticleReadCalendarCountView)
	rg.GET("articles/text", app.FullTextSearchView)
	rg.GET("articles/digg", app.ArticleDiggView)
	rg.PUT("articles", middleware.CheckAdminToken(), app.ArticleUpdateView)
	rg.DELETE("articles", middleware.CheckAdminToken(), app.ArticleRemoveView)
	rg.DELETE("articles/collects", middleware.CheckAuthToken(), app.ArticleDeleteCollectView)
}
