/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:23:21
 */

package user

import "medium-server-go/framework/db"

type User struct {
	db.Model         // 通用参数
	Phone     string `json:"phone" gorm:"unique_index, not null;"` // 手机号
	Ip        string `json:"ip"`                                   // ip 地址
	UserAgent string `json:"userAgent"`                            // 用户 ua
}
