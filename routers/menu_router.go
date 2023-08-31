package routers

import (
	"blog_server/api"
)

func (rg RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	rg.POST("menus", app.MenuCreateView)
}
