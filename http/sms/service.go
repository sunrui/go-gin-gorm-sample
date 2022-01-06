/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:41:07
 */

package sms

import (
	"fmt"
	"medium-server-go/common/app"
)

func FindByPhone(phone string) *Code {
	first := app.Db.First(&Code{}, "phone = ?", phone)
	fmt.Println(first)

	return nil
}
