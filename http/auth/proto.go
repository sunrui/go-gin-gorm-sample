/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package auth

type LoginByPhoneReq struct {
	Phone   string `json:"phone" binding:"required" validate:"min=11,max=11"`
	SmsCode string `json:"smsCode" binding:"required" validate:"min=6,max=6"`
}

//PhoneNotExist bool   `json:"phoneNotExist"`

type LoginByPhoneRes struct {
	UserId string `json:"userId"`
}
