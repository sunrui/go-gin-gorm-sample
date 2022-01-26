/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/26 13:40:26
 */

package provider

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/result"
)

// 授权中间件
func authMiddleware(ctx *gin.Context) {
	_, err := Token.GetTokenEntity(ctx)
	if err != nil {
		app.Response(ctx, result.NoAuth)
	}
}

// 管理中间件
func adminMiddleware(ctx *gin.Context) {
	_, err := Token.GetTokenEntity(ctx)
	if err != nil {
		app.Response(ctx, result.NoAuth)
	}
}

// 初始化
func init() {
	app.AuthMiddleware = authMiddleware
	app.AdminMiddleware = adminMiddleware
}
