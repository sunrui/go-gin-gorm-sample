package main

import (
	"medium-server-go/common/gin"
	"medium-server-go/controller/auth"
)

func main() {
	starter := gin.New()

	auth.RegisterHandler(starter)

	starter.Run(8080)
}
