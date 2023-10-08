package routers

import (
	"blog_server/api"
	"blog_server/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (rg RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	rg.Use(sessions.Sessions("sessionid", store)) //
	rg.POST("email_login", app.EmailLoginView)
	rg.POST("logout", middleware.CheckAuthToken(), app.UserLogoutView)
	rg.POST("user_bind_email", middleware.CheckAuthToken(), app.UserBindEmailView)
	rg.POST("login", app.QQLoginView)
	rg.POST("users", middleware.CheckAdminToken(), app.UserCreateView)
	rg.GET("users", middleware.CheckAuthToken(), app.UserListView)
	rg.GET("user_info", middleware.CheckAuthToken(), app.UserInfoView)
	rg.PUT("users", middleware.CheckAdminToken(), app.UserUpdateView)
	rg.PUT("user_password", middleware.CheckAuthToken(), app.UserUpdatePassword)
	rg.PUT("user_info", middleware.CheckAuthToken(), app.UserUpdateNickName)
	rg.DELETE("users", middleware.CheckAdminToken(), app.UserRemoveListView)
}
