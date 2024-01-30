package main

import (
	"log"

	"github.com/ankur-toko/quick-links/webserver"
)

func main() {
	err := webserver.Start()
	if err != nil {
		log.Print(err)
	}
}
