/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/06 20:30:06
 */

package sms

import (
	"gorm.io/gorm"
)

type CodeType string

const (
	LOGIN = "LOGIN"
)

type Code struct {
	gorm.Model
	Phone    string   `json:"phone" gorm:"column:phone"`
	CodeType CodeType `json:"codeType"`
	Code     string   `json:"code"`
	Day      string   `json:"day"`
	Ip       string   `json:"ip"`
}
