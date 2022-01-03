/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

const (
	LOGIN = iota
)

type PostCodeReq struct {
	Phone    string `json:"phone" binding:"required" validate:"min=11,max=11"`
	CodeType string `json:"codeType" binding:"required" validate:"oneof=LOGIN"`
}

type PostCodeRes struct {
	SendLimit bool `json:"sendLimit"`
}

type PostVerifyReq struct {
	Phone    string `json:"phone" binding:"required" validate:"min=11,max=11"`
	CodeType string `json:"codeType" binding:"required" validate:"oneof=LOGIN"`
	Code     string `json:"code" binding:"required" validate:"min=11,max=11"`
}

type PostVerifyRes struct {
	SendExpired bool `json:"sendExpired"`
	NotMatch    bool `json:"notMatch"`
}
