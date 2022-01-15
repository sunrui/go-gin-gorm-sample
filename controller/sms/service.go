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

// 初始化
func init() {
	// 创建表
	err := db.Mysql.AutoMigrate(&Code{})
	if err != nil {
		panic(err.Error())
	}
}

// 获取当天日期，如 2022-01-01
func getNowDate() string {
	now := time.Now()
	date := fmt.Sprintf("%4d-%02d-%02d", now.Year(), now.Month(), now.Day())

	return date
}

// 创建 6 位数字
func randomCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// 存储数据
func saveCode(code *Code) {
	db.Mysql.Save(code)
}

// 根据手机号和验证码类型查找
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

// 获取当天验证码发送次数
func countByPhoneAndDate(phone string, date string) int64 {
	var count int64

	query := db.Mysql.Model(&Code{}).Where("phone = ? AND DATE(created_at) = ?", phone, date).Count(&count)
	if query.Error != nil {
		panic(query.Error.Error())
	}

	return count
}
