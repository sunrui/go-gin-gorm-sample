/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/16 05:44:16
 */

package sms

import (
	"medium-server-go/framework/db"
	"medium-server-go/service/enum"
)

// 验证码对象
type Code struct {
	db.Model                // 通用参数
	Phone     string        `json:"phone"`     // 手机号
	CodeType  enum.CodeType `json:"codeType"`  // 短信类型
	Code      string        `json:"code"`      // 验证码
	Ip        string        `json:"ip"`        // ip 地址
	UserAgent string        `json:"userAgent"` // 用户 ua
	Success   bool          `json:"success"`   // 是否发送成功
	Comment   string        `json:"comment"`   // 备注
}
