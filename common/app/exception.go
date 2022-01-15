/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:51:03
 */

package app

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/common/config"
	"medium-server-go/common/result"
)

// 异常捕获对象
func exceptionHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 为了更好的调试，在开发环境中输出系统错误。
		if !config.IsDebugMode() {
			// 捕获对象，全部抛出可以使用 panic 方法。
			defer func() {
				if err := recover(); err != nil {
					dataMap := make(map[string]interface{})

					// 判断是否抛出了 result 对象
					res, ok := err.(*result.Result)
					if ok {
						dataMap["error"] = res.Data
					} else {
						dataMap["error"] = err
					}

					Response(ctx, result.InternalError.WithData(dataMap))
				}
			}()
		}

		ctx.Next()
		handlerFunc(ctx)
	}
}