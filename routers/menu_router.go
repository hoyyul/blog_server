package routers

import (
	"blog_server/api"
)

func (rg RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	rg.POST("menu", app.MenuCreateView)
	rg.GET("menu/:id", app.MenuGetView)
	rg.GET("menu", app.MenuReadListView)
	rg.GET("menu_name", app.MenuReadNameList)
	rg.PUT("menu/:id", app.MenuUpdateView)
	rg.DELETE("menu", app.MenuRemoveView)
}
