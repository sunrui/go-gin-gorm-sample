/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package sms

import "medium-server-go/framework/app"

// 获取短信路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName: "/sms",
		RoleType:  app.RolePublic,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   "POST",
				RelativePath: "/code",
				HandlerFunc:  postCode,
			}, {
				HttpMethod:   "POST",
				RelativePath: "/verify",
				HandlerFunc:  postVerify,
			},
		},
	}
}
