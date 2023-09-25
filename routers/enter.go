package routers

import (
	"blog_server/global"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //block some logs in this env
	router := gin.Default()
	router.StaticFS("uploads", http.Dir("uploads"))
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}

	//Configure API
	routerGroupApp.SettingRouter()
	routerGroupApp.ImageRouter()
	routerGroupApp.AdvertiseRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.TagRouter()
	routerGroupApp.MessageRouter()
	routerGroupApp.ArticleRouter()
	//routerGroupApp.DiggRouter()
	routerGroupApp.CommentRouter()
	routerGroupApp.NewsRouter()
	routerGroupApp.ChatRouter()
	routerGroupApp.LogRouter()
	routerGroupApp.StatisticRouter()
	return router
}
