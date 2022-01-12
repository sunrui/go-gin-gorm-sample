/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 18:07:03
 */

package sms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/app"
	"medium-server-go/common/result"
	"medium-server-go/provider"
)

const codeLimitPerDate = 5

// 发送验证码
func postCode(ctx *gin.Context) {
	var req postCodeReq

	// 较验参数
	ok, data := app.ValidateParameter(ctx, &req)
	if !ok {
		app.Response(ctx, result.ParameterError.WithData(data))
		return
	}

	// 获取当天发送条数
	count := countByPhoneAndDate(req.Phone, getNowDate())
	if count >= codeLimitPerDate {
		app.Response(ctx, &result.RateLimit)
		return
	}

	// 创建 6 位验证码
	sixNumber := createSixNumber()
	smsProvider := provider.Sms{}

	// 调用服务发送验证码
	channel, reqId, err := smsProvider.Send(req.Phone, req.CodeType, sixNumber)

	// 存储发送记录
	saveCode(&Code{
		Phone:     req.Phone,
		CodeType:  req.CodeType,
		Code:      sixNumber,
		Ip:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Success:   err != nil,
		Comment:   fmt.Sprintf("channel = %s, reqId = %s", channel, reqId),
	})

	if err != nil {
		//db.Redis.Set("hello", "world", 5e10)
	}

	// 发送成功
	app.Response(ctx, &result.Ok)
}

// 较验验证码
func postVerify(ctx *gin.Context) {
	var req postVerifyReq

	// 较验参数
	ok, data := app.ValidateParameter(ctx, &req)
	if !ok {
		app.Response(ctx, result.ParameterError.WithData(data))
		return
	}

	// 查找用户发送记录
	code := findByPhoneAndCodeType(req.Phone, string(req.CodeType))
	if code == nil || code.CodeType != req.CodeType {
		app.Response(ctx, &result.NotFound)
		return
	}

	// 比较验证码
	if code.Code != req.Code {
		app.Response(ctx, &result.NotMatch)
		return
	}

	// 较验成功
	app.Response(ctx, &result.Ok)
}
