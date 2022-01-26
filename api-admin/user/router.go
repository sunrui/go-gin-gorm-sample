/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:47:21
 */

package user

import "medium-server-go/framework/app"

// 获取短信路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName: "/user",
		RoleType:  app.RoleAdmin,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   "POST",
				RelativePath: "/:id",
				HandlerFunc:  getUser,
			},
		},
	}
}
