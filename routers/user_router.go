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
	rg.Use(sessions.Sessions("sessionid", store))
	rg.POST("email_login", app.EmailLoginView)
	rg.POST("logout", middleware.CheckAuthToken(), app.UserLogoutView)
	rg.POST("user_bind_email", middleware.CheckAuthToken(), app.UserBindEmailView)
	rg.GET("user", middleware.CheckAuthToken(), app.UserListView)
	rg.PUT("user", middleware.CheckAdminToken(), app.UserUpdateView)
	rg.PUT("user_password", middleware.CheckAuthToken(), app.UserUpdatePassword)
	rg.DELETE("users", middleware.CheckAdminToken(), app.UserDeleteListView)
}
