package main

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/application"
	"log"
)

func main() {
	router, err := application.CreateRouter()

	if err != nil {
		log.Fatal("Error configurating the router", err)
	}

	router.Run()
}
