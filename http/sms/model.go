/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/06 20:30:06
 */

package sms

import "medium-server-go/common/db"

type CodeType string

const (
	LOGIN = "LOGIN"
)

type Code struct {
	db.Model
	Phone    string   `json:"phone"`
	CodeType CodeType `json:"codeType"`
	Code     string   `json:"code"`
	Day      string   `json:"day"`
	Ip       string   `json:"ip"`
}
