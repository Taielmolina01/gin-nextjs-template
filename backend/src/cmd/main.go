package main

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/application"
)

func main() {
	router := application.CreateRouter()

	router.Run()
}