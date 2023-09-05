package routers

import (
	"blog_server/api"
	"blog_server/middleware"
)

func (rg RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	rg.POST("email_login", app.EmailLoginView)
	rg.POST("logout", middleware.CheckAuthToken(), app.UserLogoutView)
	rg.GET("user", middleware.CheckAuthToken(), app.UserListView)
	rg.PUT("user", middleware.CheckAdminToken(), app.UserUpdateView)
	rg.PUT("user_password", middleware.CheckAuthToken(), app.UserUpdatePassword)
	rg.DELETE("users", middleware.CheckAdminToken(), app.UserDeleteListView)
}
