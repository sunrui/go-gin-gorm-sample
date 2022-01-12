/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/02
 */
package auth

import (
	"medium-server-go/common/app"
)

// 获取授权路由对象
func GetRouter() app.Router {
	return app.Router{
		GroupName: "/auth",
		NeedAuth:  false,
		RouterPaths: []app.RouterPath{
			{
				HttpMethod:   "POST",
				RelativePath: "/login/phone",
				HandlerFunc:  postLoginByPhone,
			}, {
				HttpMethod:   "POST",
				RelativePath: "/login/wechat",
				HandlerFunc:  postLoginByWechat,
			},
		},
	}
}
