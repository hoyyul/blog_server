package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	rg.POST("menus", middleware.CheckAdminToken(), app.MenuUploadView)
	rg.GET("menus/:id", app.MenuDetailView)
	rg.GET("menus", app.MenuListView)
	rg.GET("menu_names", app.MenuNameList)
	rg.GET("menus/detail", app.MenuDetailByPathView)
	rg.PUT("menus/:id", middleware.CheckAdminToken(), app.MenuUpdateView)
	rg.DELETE("menus", middleware.CheckAdminToken(), app.MenuRemoveListView)
}
