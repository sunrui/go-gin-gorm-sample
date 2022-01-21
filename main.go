/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */
package main

import (
	"medium-server-go/api-public/auth"
	"medium-server-go/api-public/sms"
	"medium-server-go/api-user/user"
	"medium-server-go/framework/app"
)

func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	for _, router := range []app.Router{
		sms.GetRouter(),
		auth.GetRouter(),
		user.GetRouter(),
	} {
		server.RegisterRouter(router)
	}

	// 启动服务
	server.Run(8080)
}
