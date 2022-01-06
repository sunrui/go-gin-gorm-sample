/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 18:07:03
 */

package sms

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/common/app"
	"medium-server-go/common/result"
	"net/http"
)

func PostCode(ctx *gin.Context) {
	var req PostCodeReq

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		app.Response(ctx, errNo)
		return
	}

	count := CountByPhoneAndToday(req.Phone)
	if count > 5 {
		dataMap := make(map[string]int)
		dataMap["count"] = count

		app.Response(ctx, result.RateLimit.WithData(dataMap))
		return
	}

	code := FindByPhone(req.Phone)
	if code == nil {
		dataMap := make(map[string]string)
		dataMap["phone"] = req.Phone

		app.Response(ctx, result.NotFound.WithData(dataMap))
		return
	}

	if req.Phone != "15068860507" {
		app.Response(ctx, &result.NotFound)
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok)
}

func PostVerify(ctx *gin.Context) {
	var req PostVerifyReq

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	if req.Phone != "15068860507" {
		ctx.JSON(http.StatusOK,
			result.Ok)
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok)
}
