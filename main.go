package main

import (
	"medium-server-go/common/starter"
	"medium-server-go/controller/auth"
)

func main() {
	app := starter.New()

	app.RegisterRouter(auth.GetRouter())

	app.Run(8080)
}
