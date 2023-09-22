package routers

import "blog_server/api"

func (rg RouterGroup) NewsRouter() {
	app := api.ApiGroupApp.NewsApi
	rg.POST("news", app.NewsReadListView)
}
