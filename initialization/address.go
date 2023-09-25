package initialization

import (
	"blog_server/global"
	"log"

	geoip2db "github.com/cc14514/go-geoip2-db"
)

func InitAddrDB() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		log.Fatal("Failed to get ip address", err)
	}
	global.AddrDB = db
}
