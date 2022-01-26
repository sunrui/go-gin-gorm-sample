/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/25 22:29:25
 */

package api_public

import (
	"medium-server-go/api-public/auth"
	"medium-server-go/api-public/sms"
	"medium-server-go/framework/app"
)

// 注册路由
func Register(server *app.Server) {
	for _, router := range []app.Router{
		sms.GetRouter(),
		auth.GetRouter(),
	} {
		server.RegisterRouter(router)
	}
}
