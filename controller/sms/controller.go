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

func postCode(ctx *gin.Context) {
	var req postCodeReq

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		app.Response(ctx, errNo)
		return
	}

	count := countByPhoneAndDate(req.Phone, getNowDate())
	if count >= codeLimitPerDate {
		app.Response(ctx, &result.RateLimit)
		return
	}

	sixNumber := createSixNumber()
	smsProvider := provider.Sms{}

	channel, reqId, err := smsProvider.Send(req.Phone, req.CodeType, sixNumber)

	createCode(&Code{
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

	app.Response(ctx, &result.Ok)
}

func postVerify(ctx *gin.Context) {
	var req postVerifyReq

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		app.Response(ctx, errNo)
		return
	}

	code := findByPhoneAndCodeType(req.Phone, string(req.CodeType))
	if code == nil || code.CodeType != req.CodeType {
		app.Response(ctx, &result.NotFound)
		return
	}

	if code.Code != req.Code {
		app.Response(ctx, &result.NotMatch)
		return
	}

	app.Response(ctx, &result.Ok)
}
