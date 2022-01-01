package auth

import "medium-server-go/common/errno"

var (
	UserNotExist  = errno.MakeErrNo("UserNotExist", "用户不存在")
	PasswordError = errno.MakeErrNo("PasswordError", "密码不正确")
)
