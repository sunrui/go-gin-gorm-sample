/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 18:07:03
 */

package sms

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/app"
	"medium-server-go/common/result"
	"medium-server-go/provider"
)

// 发送验证码
func postCode(ctx *gin.Context) {
	var req postCodeReq

	// 较验参数
	data, ok := app.ValidateParameter(ctx, &req)
	if !ok {
		app.Response(ctx, result.ParameterError.WithData(data))
		return
	}

	// 获取当天发送条数，判断是否超出最大条数限制
	count := countByPhoneAndDate(req.Phone, getNowDate())
	if count >= 5 {
		app.Response(ctx, &result.RateLimit)
		return
	}

	// 创建 6 位验证码
	sixNumber := createSixNumber()
	smsProvider := provider.Sms{}

	// 调用服务发送验证码
	channel, reqId, err := smsProvider.Send(req.Phone, req.CodeType, sixNumber)

	// 备注对象
	type comment struct {
		Chanel string `json:"chanel"`
		ReqId  string `json:"reqId"`
	}

	// 将备注对象转 json
	marshal, _ := json.Marshal(comment{
		Chanel: channel,
		ReqId:  reqId,
	})

	// 存储发送记录
	saveCode(&Code{
		Phone:     req.Phone,
		CodeType:  req.CodeType,
		Code:      sixNumber,
		Ip:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Success:   err == nil,
		Comment:   string(marshal),
	})

	// 发送验证码失败
	if err != nil {
		app.Response(ctx, result.InternalError.WithData(err))
		return
	}

	// 将验证码缓存到 redis 中
	cache := Cache{
		Phone:    req.Phone,
		CodeType: req.CodeType,
	}
	cache.Save(CodeCache{
		Code:      sixNumber,
		ErrVerify: 0,
	})

	// 发送成功
	app.Response(ctx, &result.Ok)
}

// 较验验证码
func postVerify(ctx *gin.Context) {
	var req postVerifyReq

	// 较验参数
	data, ok := app.ValidateParameter(ctx, &req)
	if !ok {
		app.Response(ctx, result.ParameterError.WithData(data))
		return
	}

	// 缓存对象
	cache := Cache{
		Phone:    req.Phone,
		CodeType: req.CodeType,
	}

	// 获取缓存数据
	codeCache := cache.GetValue()
	if codeCache == nil {
		app.Response(ctx, &result.NotFound)
	}

	// 如果验证码较验错误
	if codeCache.Code != req.Code {
		// 增加缓存引用记数
		codeCache.ErrVerify += 1

		// 如果已经较验出错 5 次，移除现有验证码
		if codeCache.ErrVerify == 5 {
			cache.Del()
		} else {
			// 更新出错较验次数
			cache.Save(*codeCache)
		}

		app.Response(ctx, &result.NotMatch)
		return
	}

	// 较验成功
	app.Response(ctx, &result.Ok)
}
