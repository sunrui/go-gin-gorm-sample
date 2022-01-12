/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:41:07
 */

package sms

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"medium-server-go/common/db"
	"time"
)

func init() {
	err := db.Mysql.AutoMigrate(&Code{})
	if err != nil {
		panic(err.Error())
	}
}

func getNowDate() string {
	now := time.Now()
	date := fmt.Sprintf("%4d-%02d-%02d", now.Year(), now.Month(), now.Day())

	return date
}

func createSixNumber() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func createCode(code *Code) {
	db.Mysql.Save(code)
}

func findByPhoneAndCodeType(phone string, codeType string) *Code {
	var code Code
	query := db.Mysql.Where("phone = ? AND code_type = ? AND date(created_at) = ?", phone, codeType, getNowDate()).Last(&code)
	if query.Error == gorm.ErrRecordNotFound {
		return nil
	}

	if query.Error != nil {
		panic(query.Error.Error())
	}

	return &code
}

func countByPhoneAndDate(phone string, date string) int64 {
	var count int64

	query := db.Mysql.Model(&Code{}).Where("phone = ? AND DATE(created_at) = ?", phone, date).Count(&count)
	if query.Error != nil {
		panic(query.Error.Error())
	}

	return count
}
