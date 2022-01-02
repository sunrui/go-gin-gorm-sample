/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date:2021/12/31
 */
package main

import (
	"medium-server-go/common/app"
	"medium-server-go/controller/auth"
)

func main() {
	inst := app.New()

	inst.RegisterRouter(auth.GetRouter())

	inst.Run(8080)
}
