/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:41:07
 */

package sms

import (
	"fmt"
	"medium-server-go/common/db"
	"time"
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

func countByPhoneAndDay(phone string, day string) int64 {
	var count int64
	query := db.Default.Where("phone = ? AND day = ?", phone, day).Find(&Code{}).Count(&count)
	if query.Error != nil {
		panic(query.Error.Error())
	}

	//db.Default.Raw(fmt.Sprintf("SELECT count(*) FROM code WHERE date(created_at) = %"))

	return count
}
