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

type Code struct {
	db.Model
	Phone     string        `json:"phone"`
	CodeType  enum.CodeType `json:"codeType"`
	Code      string        `json:"code"`
	Ip        string        `json:"ip"`
	UserAgent string        `json:"userAgent"`
	Success   bool          `json:"success"`
	Comment   string        `json:"comment"`
}
