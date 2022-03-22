package main

import (
	"fmt"
	"log"
	"net/url"

)



func main() {

	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal("failed error ")
	}

	switch u.Scheme {
	case "unix":
		fmt.Println(u.Path)
	}
}
