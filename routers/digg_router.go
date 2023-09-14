package routers

import "blog_server/api"

func (rg RouterGroup) DiggRouter() {
	app := api.ApiGroupApp.DiggApi
	rg.POST("digg/article", app.DiggArticleView)
}
