package routers

import (
	"blog_server/api"
)

func (rg RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi
	rg.GET("chat_groups", app.ChatGroupView)
	rg.GET("chat_groups_records", app.ChatListView)
}
