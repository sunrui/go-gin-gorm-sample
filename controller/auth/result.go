package auth

import (
	"medium-server-go/common/result"
)

var (
	UserNotExist  = result.MakeResult("UserNotExist", "用户不存在")
	PasswordError = result.MakeResult("PasswordError", "密码不正确")
)
