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

func catch(ctx *gin.Context) {
	if err := recover(); err != nil {
		dataMap := make(map[string]interface{})

		res, ok := err.(*result.Result)
		if ok {
			dataMap["error"] = res.Data
		} else {
			dataMap["error"] = err
		}

		ret := result.InternalError.WithData(dataMap)
		Response(ctx, result.InternalError.WithData(ret))
	}
}

func catchHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !config.IsDebugMode() {
			defer catch(ctx)
		}

		ctx.Next()
		handlerFunc(ctx)
	}
}
