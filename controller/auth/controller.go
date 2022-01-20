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
	"medium-server-go/enum"
	"medium-server-go/provider"
	"net/http"
)

// 手机号码登录
func postLoginByPhone(ctx *gin.Context) {
	var req loginByPhoneReq

	// 较验参数
	errData, err := app.ValidateParameter(ctx, &req)
	if err != nil {
		app.Response(ctx, result.ParameterError.WithData(errData))
		return
	}

	// 短信缓存对象
	smsCache := sms.Cache{
		Phone:    req.Phone,
		CodeType: enum.Code.Login,
	}

	// 获取缓存数据
	if !smsCache.Exists() {
		app.Response(ctx, result.NotFound)
		return
	}

	// 较验验证码
	if !smsCache.Verify(req.Code) {
		app.Response(ctx, result.NotMatch)
		return
	}

	// 移除验证码
	smsCache.Del()

	token, err := provider.Token.Encode(provider.TokenEntity{
		UserId: "userId",
	})
	if err != nil {
		return
	}

	ctx.SetCookie("token", token, 30*24*60*60,
		"/", "localhost", false, true)

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(loginByPhoneRes{
			UserId: req.Phone,
		}))
}

// 微信登录
func postLoginByWechat(ctx *gin.Context) {
	var req loginByPhoneReq

	// 较验参数
	errData, err := app.ValidateParameter(ctx, &req)
	if err != nil {
		app.Response(ctx, result.ParameterError.WithData(errData))
		return
	}
}

// 获取令牌
func getToken(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		app.Response(ctx, result.NotFound)
		return
	}

	// 获取用户令牌
	tokenEntity, err := provider.Token.Decode(token)
	if err != nil {
		panic(err)
		return
	}

	app.Response(ctx, result.Ok.WithData(tokenEntity))
}

// 登出
func postLogout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")
	if err != nil {
		app.Response(ctx, result.NotFound)
		return
	}

	// 设置令牌过期
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	app.Response(ctx, result.Ok)
}
