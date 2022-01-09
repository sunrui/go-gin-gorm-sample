/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:50:03
 */

package app

import (
	"github.com/gin-gonic/gin"
)

type RouterPath struct {
	HttpMethod   string
	RelativePath string
	HandlerFunc  gin.HandlerFunc
}

type Router struct {
	GroupName   string
	NeedAuth    bool
	RouterPaths []RouterPath
}

func (app *Server) RegisterRouter(router Router) {
	groupRouter := app.engine.Group(router.GroupName)

	if router.NeedAuth {
		groupRouter.Use(authMiddleware)
	}

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
