package main

import (
	_ "blog_server/docs"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/routers"
	"blog_server/utils"
)

// @title blog_server API Documentation
// @version 1.0
// @description blog_server API Documentation
// @host 127.0.0.01:8080
// @BasePath /

func main() {
	// Initialize configuration settings
	initialization.InitConf()

	// Initialize logger
	global.Logger = initialization.InitLogger()

	// Connect database
	global.DB = initialization.InitGorm()

	// Connect redis
	global.Redis = initialization.InitRedis()

	// Connect elasticSearch
	global.ESClient = initialization.EsConnect()

	// Initialize router
	router := routers.InitRouter()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.RunOption(option)
		return
	}

	addr := global.Config.System.Addr()
	utils.PrintSystem()
	router.Run(addr)

	err := router.Run(addr)
	if err != nil {
		global.Logger.Fatalf(err.Error())
	}
}
