package routers

import "blog_server/api"

func (rg RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	rg.POST("tag", app.TagCreateView)
	rg.GET("tag", app.TagReadListView)
	rg.PUT("tag/:id", app.TagUpdateView)
	rg.DELETE("tag", app.TagDeleteListView)
}
