/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */
package main

import (
	"medium-server-go/common/app"
	"medium-server-go/common/db"
	"medium-server-go/http/auth"
	"medium-server-go/http/sms"
)

func main() {
	db.Test()

	return

	inst := app.New()

	inst.RegisterRouter(sms.GetRouter())
	inst.RegisterRouter(auth.GetRouter())

	inst.Run(8080)
}
