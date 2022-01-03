/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/02
 */
package auth

import (
	"medium-server-go/common/app"
)

func GetRouter() app.Router {
	var router app.Router

	router.GroupName = "/auth"
	router.NeedAuth = true
	router.RouterPaths = []app.RouterPath{
		{
			HttpMethod:   "POST",
			RelativePath: "/login/phone",
			HandlerFunc:  PostLoginByPhone,
		}, {
			HttpMethod:   "POST",
			RelativePath: "/login/wechat",
			HandlerFunc:  PostLoginByWechat,
		},
	}

	return router
}
