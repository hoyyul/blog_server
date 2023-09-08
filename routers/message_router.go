package routers

import (
	"blog_server/api"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("messages", app.MessageCreateView)
}
