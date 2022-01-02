// Copyright 2022 honeysense  All rights reserved.
// Author: sunrui, smallrui@foxmail.com
// Date: 2021.1.2 22:31
//
package auth

import (
	"medium-server-go/common/result"
)

var (
	UserNotExist  = result.MakeResult("UserNotExist", "用户不存在")
	PasswordError = result.MakeResult("PasswordError", "密码不正确")
)
