/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package auth

// 手机号码登录请求
type loginByPhoneReq struct {
	Phone string `json:"phone" binding:"required" validate:"len=11"` // 手机号
	Code  string `json:"code" binding:"required" validate:"len=6"`   // 验证码
}

// 手机号码登录结果
type loginByPhoneRes struct {
	UserId string `json:"userId"` // 用户 id
}
