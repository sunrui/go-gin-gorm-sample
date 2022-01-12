/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package auth

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/common/app"
	"medium-server-go/common/result"
	"net/http"
)

// 手机号码登录
func postLoginByPhone(ctx *gin.Context) {
	var req loginByPhoneReq

	// 较验参数
	haveError, dataMap := app.ValidateParameter(ctx, &req)
	if haveError {
		app.Response(ctx, result.ParameterError.WithData(dataMap))
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(loginByPhoneRes{
			UserId: req.Phone,
		}))
}

// 微信登录
func postLoginByWechat(ctx *gin.Context) {
	var req loginByPhoneReq

	// 较验参数
	haveError, dataMap := app.ValidateParameter(ctx, &req)
	if haveError {
		app.Response(ctx, result.ParameterError.WithData(dataMap))
		return
	}
}
