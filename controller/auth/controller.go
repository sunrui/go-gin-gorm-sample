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

	errNo := app.ValidateParameter(ctx, &req)
	if errNo != nil {
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(loginByPhoneRes{
			UserId: req.Phone,
		}))
}

// 微信登录
func postLoginByWechat(ctx *gin.Context) {

}
