package routers

import (
	"blog_server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //block some logs in this env
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.String(200, "xxx")
	})
	return router
}
