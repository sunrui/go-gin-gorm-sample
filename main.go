package main

import (
	"medium-server-go/common/gin"
	"medium-server-go/controller/auth"
)

func main() {
	app := gin.New()

	app.RegisterRouter(auth.GetRouter())

	app.Run(8080)
}
