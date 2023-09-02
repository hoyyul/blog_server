package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string

	flag.StringVar(&user, "u", "", "create user") // if input value == "", usage will be printed
	flag.Parse()

	fmt.Println(user)
}
