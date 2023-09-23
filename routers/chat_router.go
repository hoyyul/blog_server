package routers

import (
	"blog_server/api"
)

func (rg RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi
	rg.GET("chat_group", app.ChatGroupView)
}
