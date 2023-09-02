package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	rg.POST("email_login", app.EmailLoginView)
	rg.GET("users", middleware.CheckAuthToken(), app.UserListView)
}
