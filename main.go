package main

import (
	"blog_server/global"
	"blog_server/initialization"
)

func main() {
	// Initialize configuration
	initialization.InitConf()

	//Initialize logger
	global.Logger = initialization.InitLogger()

	//Connect database
	global.DB = initialization.InitGorm()

}
