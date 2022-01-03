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
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	if req.Phone != "15068860507" {
		ctx.JSON(http.StatusOK,
			result.Ok.WithData(PostCodeRes{}))
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(PostCodeRes{}))
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
			result.Ok.WithData(PostVerifyRes{}))
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(PostVerifyRes{}))
}
