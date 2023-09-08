package routers

import "blog_server/api"

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tags", app.TagCreateView)
	router.GET("tags", app.TagReadListView)
	router.PUT("tags/:id", app.TagUpdateView)
	router.DELETE("tags", app.TagDeleteListView)
}
