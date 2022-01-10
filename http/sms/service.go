/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:41:07
 */

package sms

import (
	"medium-server-go/common/db"
)

func init() {
	err := db.Default.AutoMigrate(&Code{})
	if err != nil {
		panic(err.Error())
	}
}

func createCode(code *Code) {
	db.Default.Save(code)
}

func findByPhoneAndCodeType(phone string, codeType string) *Code {
	var code *Code
	query := db.Default.First(code, "phone = ? AND codeType = ?", phone, codeType)
	if query.Error != nil {
		panic(query.Error.Error())
	}

	return code
}

func countByPhoneAndDate(phone string, date string) int64 {
	var count int64

	query := db.Default.Where("phone = ? AND DATE(created_at) = ?", phone, date).Find(&Code{}).Count(&count)
	if query.Error != nil {
		panic(query.Error.Error())
	}

	return count
}
