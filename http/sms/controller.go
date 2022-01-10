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
	"time"
)

const codeLimitPerDate = 5

func postCode(ctx *gin.Context) {
	var req postCodeReq

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		app.Response(ctx, errNo)
		return
	}

	now := time.Now()
	date := fmt.Sprintf("%4d-%02d-%02d", now.Year(), now.Month(), now.Day())

	count := countByPhoneAndDate(req.Phone, date)
	if count >= codeLimitPerDate {
		app.Response(ctx, &result.RateLimit)
		return
	}

	createCode(&Code{
		Phone:    req.Phone,
		CodeType: req.CodeType,
		Code:     "123456",
		Ip:       ctx.ClientIP(),
	})

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
	if code == nil {
		app.Response(ctx, &result.NotFound)
		return
	}

	if req.Phone != "15068860507" {
		app.Response(ctx, &result.Ok)
		return
	}

	app.Response(ctx, &result.Ok)
}
