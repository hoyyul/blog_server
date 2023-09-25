package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	rg.POST("menu", middleware.CheckAdminToken(), app.MenuCreateView)
	rg.GET("menu/:id", app.MenuGetView)
	rg.GET("menu", app.MenuReadListView)
	rg.GET("menu_name", app.MenuReadNameList)
	rg.PUT("menu/:id", middleware.CheckAdminToken(), app.MenuUpdateView)
	rg.DELETE("menu", middleware.CheckAdminToken(), app.MenuRemoveView)
}
