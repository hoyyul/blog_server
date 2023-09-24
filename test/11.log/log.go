package main

import (
	"blog_server/global"
	"blog_server/initialization"
	"blog_server/plugins/log_stash"
	"fmt"
)

func main() {
	initialization.InitConf()
	global.Logger = initialization.InitLogger()

	global.DB = initialization.InitGorm()
	log := log_stash.New("192.168.100.158", "xxxx")

	log.Error(fmt.Sprintf("%s hello", "hy"))
}
