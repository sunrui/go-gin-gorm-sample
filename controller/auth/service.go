// Copyright 2022 honeysense  All rights reserved.
// Author: sunrui, smallrui@foxmail.com
// Date: 2021.1.1 23:35
//
package auth

import (
	"github.com/gin-gonic/gin"
	starter "medium-server-go/common/gin"
	"medium-server-go/common/result"
	"net/http"
)

func LoginByPhone(ctx *gin.Context) {
	var req LoginByPhoneReq

	errNo := starter.ValidateParameter(ctx, &req)
	if errNo != nil {
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	if req.Phone != "15068860507" {
		ctx.JSON(http.StatusOK,
			result.Ok.WithData(LoginByPhoneRes{
				PhoneNotExist: true,
			}))
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(LoginByPhoneRes{
			UserId: req.Phone,
		}))
}

func LoginByWechat(ctx *gin.Context) {

}
