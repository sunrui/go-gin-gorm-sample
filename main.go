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

	server.RegisterRouter(sms.GetRouter())
	server.RegisterRouter(auth.GetRouter())

	server.Run(8080)
}
