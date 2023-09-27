package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	rg.POST("menus", middleware.CheckAdminToken(), app.MenuCreateView)
	rg.GET("menus/:id", app.MenuGetView)
	rg.GET("menus", app.MenuReadListView)
	rg.GET("menu_names", app.MenuReadNameList)
	rg.PUT("menus/:id", middleware.CheckAdminToken(), app.MenuUpdateView)
	rg.DELETE("menus", middleware.CheckAdminToken(), app.MenuRemoveView)
}
