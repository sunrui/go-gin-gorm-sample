// Copyright 2022 honeysense  All rights reserved.
// Author: sunrui, smallrui@foxmail.com
// Date: 2021.1.1 23:25
//
package auth

type LoginByPhoneReq struct {
	Phone   string `json:"phone" binding:"required" validate:"min=11,max=11"`
	SmsCode string `json:"smsCode" binding:"required" validate:"min=6,max=6"`
}

type LoginByPhoneRes struct {
	PhoneNotExist bool   `json:"phoneNotExist"`
	UserId        string `json:"userId"`
}
