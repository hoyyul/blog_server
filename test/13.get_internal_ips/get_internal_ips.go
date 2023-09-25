package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			logrus.Error(err)
			continue
		}
		//fmt.Println(i2.Name, address)
		// filter ipv6
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ip4 := ipNet.IP.To4()
			if ip4 == nil {
				continue
			}
			fmt.Println(ip4)
		}
	}

}
