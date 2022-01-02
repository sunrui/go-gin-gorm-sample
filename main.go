package main

import (
	"medium-server-go/common/gin"
	"medium-server-go/controller/auth"
)

func main() {
	inst := app.New()

	inst.RegisterRouter(auth.GetRouter())

	inst.Run(8080)
}
