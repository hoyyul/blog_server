package utils

import (
	"blog_server/global"
)

func PrintSystem() {

	ip := global.Config.System.Host
	port := global.Config.System.Port

	if ip == "0.0.0.0" {
		ipList := GetIPList()
		for _, ip := range ipList {
			global.Logger.Infof("gvb_server is running on http://%s:%d/api", ip, port)
		}
	} else {
		global.Logger.Infof("gvb_server is running on http://%s:%d/api", ip, port)
	}

}
