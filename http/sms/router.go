/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:59:03
 */

package sms

import "medium-server-go/common/app"

func GetRouter() app.Router {
	var router app.Router

	router.GroupName = "/sms"
	router.NeedAuth = true
	router.RouterPaths = []app.RouterPath{
		{
			HttpMethod:   "POST",
			RelativePath: "/code",
			HandlerFunc:  PostCode,
		}, {
			HttpMethod:   "POST",
			RelativePath: "/verify",
			HandlerFunc:  PostVerify,
		},
	}

	return router
}