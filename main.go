package main

import (
	_ "blog_server/docs"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/routers"
)

// @title blog_server API Documentation
// @version 1.0
// @description blog_server API Documentation
// @host 127.0.0.01:8080
// @BasePath /

func main() {
	// Initialize configuration settings
	initialization.InitConf()

	//Initialize logger
	global.Logger = initialization.InitLogger()

	//Connect database
	global.DB = initialization.InitGorm()

	//Initialize router
	router := routers.InitRouter()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.RunOption(option)
	}

	addr := global.Config.System.Addr()
	global.Logger.Infof("Server is running on %s", addr)
	router.Run(addr)
}
