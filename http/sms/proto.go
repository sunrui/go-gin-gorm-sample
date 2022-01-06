/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

type PostCodeReq struct {
	Phone    string   `json:"phone" binding:"required" validate:"min=11,max=11"`
	CodeType CodeType `json:"codeType" binding:"required" validate:"oneof=LOGIN"`
}

//var (
//	SendLimit = result.MakeResult("SendLimit", "发送限制")
//)

type PostVerifyReq struct {
	Phone    string   `json:"phone" binding:"required" validate:"min=11,max=11"`
	CodeType CodeType `json:"codeType" binding:"required" validate:"oneof=LOGIN"`
	Code     string   `json:"code" binding:"required" validate:"min=11,max=11"`
}

//
//var (
//	SendExpired = result.MakeResult("SendExpired", "发送已过期")
//	NotMatch    = result.MakeResult("NotMatch", "验证码不匹配")
//)
