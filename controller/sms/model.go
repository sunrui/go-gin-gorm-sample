/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/06 20:30:06
 */

package sms

import (
	"medium-server-go/common/db"
	"medium-server-go/enum"
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
