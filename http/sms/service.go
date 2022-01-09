/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:41:07
 */

package sms

import (
	"fmt"
	"medium-server-go/common/db"
	"medium-server-go/common/result"
	"time"
)

func init() {
	err := db.Default.AutoMigrate(&Code{})
	if err != nil {
		panic(result.InternalError.WithData(err.Error()))
	}
}

func createCode(code *Code) {
	db.Default.Save(code)
}

func makeToday() string {
	now := time.Now()
	today := fmt.Sprintf("%4d%02d%02d", now.Year(), now.Month(), now.Day())
	return today
}

func findByPhone(phone string) *Code {
	first := db.Default.First(&Code{}, "phone = ?", phone)
	fmt.Println(first)

	return nil
}

func countByPhoneAndDay(phone string, day string) int {
	var code []Code
	query := db.Default.Where("phone = ? AND day = ?", phone, day).Find(&code)
	if query.Error != nil {
		panic(result.InternalError.WithData(query.Error.Error()))
	}

	fmt.Println(code)

	panic(result.InternalError.WithData(code))
	//panic(code)

	return 0
}
