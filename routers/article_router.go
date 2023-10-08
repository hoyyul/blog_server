package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	rg.POST("articles", middleware.CheckAdminToken(), app.ArticleUploadView)
	rg.POST("articles/collects", middleware.CheckAuthToken(), app.ArticleCollectView)
	rg.POST("articles/digg", app.ArticleDiggView)
	rg.GET("articles/collects", middleware.CheckAuthToken(), app.ArticleCollectListView)
	rg.GET("categorys", app.ArticleCategoryListView)
	rg.GET("articles/content/:id", app.ArticleContentByIDView)
	rg.GET("articles", app.ArticleListView)
	rg.GET("articles/:id", app.ArticleDetailView)
	rg.GET("articles/detail", app.ArticleDetailByTitleView)
	rg.GET("articles/calendar", app.ArticleCalendarCountView)
	rg.GET("articles/text", app.FullTextSearchView)
	rg.GET("article_id_title", app.ArticleIDTitleListView)
	rg.GET("articles/tags", app.ArticleTagListView)
	rg.PUT("articles", middleware.CheckAdminToken(), app.ArticleUpdateView)
	rg.DELETE("articles", middleware.CheckAdminToken(), app.ArticleRemoveView)
	rg.DELETE("articles/collects", middleware.CheckAuthToken(), app.ArticleRemoveCollectView)
}
