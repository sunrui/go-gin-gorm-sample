/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/20 23:16:20
 */

package user

import "medium-server-go/common/db"

type User struct {
	db.Model        // 通用参数
	Phone    string `json:"phone" gorm:"unique_index, not null;"` // 手机号
}
