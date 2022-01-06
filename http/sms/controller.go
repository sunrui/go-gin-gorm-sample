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
	"net/http"
)

func PostCode(ctx *gin.Context) {
	var req PostCodeReq

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	first := app.Db.First(&Code{}, "phone = ?", "15068860507")
	fmt.Println(first)

	if req.Phone != "15068860507" {
		ctx.JSON(http.StatusOK,
			result.Ok)
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
