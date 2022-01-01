package main

import (
	"medium-server-go/common/starter"
	"medium-server-go/controller/auth"
)

func main() {
	app := starter.New()
	auth.RegisterHandler(app)
	app.Run(8080)
}
