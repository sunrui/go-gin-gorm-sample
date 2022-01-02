// Copyright 2022 honeysense  All rights reserved.
// Author: sunrui, smallrui@foxmail.com
// Date: 2021.12.31 23:56
//
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
