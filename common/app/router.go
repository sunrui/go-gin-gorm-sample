/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:50:03
 */

package app

import (
	"github.com/gin-gonic/gin"
)

// 路由路径
type RouterPath struct {
	// 方法类型 GET、POST、PUT、DELETE
	HttpMethod string
	// 路径
	RelativePath string
	// 回调
	HandlerFunc gin.HandlerFunc
}

// 路由对象
type Router struct {
	// 组名
	GroupName string
	// 是否需要授权
	NeedAuth bool
	// 路由路径
	RouterPaths []RouterPath
}

// 注册路由对象
func (app *Server) RegisterRouter(router Router) {
	groupRouter := app.engine.Group(router.GroupName)

	// 如果需要授权，注册授权中间件
	if router.NeedAuth {
		groupRouter.Use(authMiddleware)
	}

	// 注册路由回调
	for _, routerPath := range router.RouterPaths {
		switch routerPath.HttpMethod {
		case "GET":
			groupRouter.GET(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "POST":
			groupRouter.POST(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "PUT":
			groupRouter.PUT(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "DELETE":
			groupRouter.DELETE(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		default:
			panic("http method not supported")
		}
	}
}
