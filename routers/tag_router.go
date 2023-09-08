package routers

import "blog_server/api"

func (rg RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	rg.POST("tags", app.TagCreateView)
	rg.GET("tags", app.TagReadListView)
	rg.PUT("tags/:id", app.TagUpdateView)
	rg.DELETE("tags", app.TagDeleteListView)
}
