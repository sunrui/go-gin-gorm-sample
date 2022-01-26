/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/25 22:29:25
 */

package api_user

import (
	"medium-server-go/api-user/user"
	"medium-server-go/framework/app"
)

// 注册路由
func Register(server *app.Server) {
	for _, router := range []app.Router{
		user.GetRouter(),
	} {
		server.RegisterRouter(router)
	}
}
