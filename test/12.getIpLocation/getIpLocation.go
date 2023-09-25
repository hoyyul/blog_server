package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("xx")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v is in %s, %s\n", ip, record.City.Names["en"], record.Country.Names["en"])
}
