/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:51:03
 */

package sms

import "medium-server-go/enum"

type postCodeReq struct {
	Phone    string        `json:"phone" binding:"required" validate:"len=11,numeric"`
	CodeType enum.CodeType `json:"codeType" binding:"required" validate:"oneof=LOGIN"`
}

type postVerifyReq struct {
	Phone    string        `json:"phone" binding:"required" validate:"len=11,numeric"`
	CodeType enum.CodeType `json:"codeType" binding:"required" validate:"oneof=LOGIN"`
	Code     string        `json:"code" binding:"required" validate:"len=6,numeric"`
}
