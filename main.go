/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */
package main

import (
	"medium-server-go/common/app"
	"medium-server-go/controller/auth"
	"medium-server-go/controller/sms"
)

func main() {
	server := app.New()

	for _, router := range []app.Router{
		sms.GetRouter(),
		auth.GetRouter(),
	} {
		server.RegisterRouter(router)
	}

	server.Run(8080)
}
