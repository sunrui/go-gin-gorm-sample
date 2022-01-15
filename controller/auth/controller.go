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
	"medium-server-go/controller/sms"
	"net/http"
)

// 手机号码登录
func postLoginByPhone(ctx *gin.Context) {
	var req loginByPhoneReq

	// 较验参数
	data, ok := app.ValidateParameter(ctx, &req)
	if !ok {
		app.Response(ctx, result.ParameterError.WithData(data))
		return
	}

	// 缓存对象
	cache := sms.Cache{
		Phone:    req.Phone,
		CodeType: "LOGIN",
	}

	// 获取缓存数据
	if !cache.Exists() {
		app.Response(ctx, &result.NotFound)
		return
	}

	// 较验验证码
	if !cache.Verify(req.Code) {
		app.Response(ctx, &result.NotMatch)
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
	data, ok := app.ValidateParameter(ctx, &req)
	if !ok {
		app.Response(ctx, result.ParameterError.WithData(data))
		return
	}

}
