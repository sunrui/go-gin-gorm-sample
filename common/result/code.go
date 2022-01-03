/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 19:03:03
 */

package result

var (
	Ok               = MakeResult("Ok", "成功")
	NoAuth           = MakeResult("NoAuth", "没有登录")
	Duplicate        = MakeResult("Duplicate", "已经存在")
	Forbidden        = MakeResult("Forbidden", "没有权限")
	NotFound         = MakeResult("NotFound", "不存在")
	NotMatch         = MakeResult("NotMatch", "不匹配")
	RateLimit        = MakeResult("ExceedLimit", "超出限制")
	LogicError       = MakeResult("LogicError", "逻辑错误")
	ParameterError   = MakeResult("ParameterError", "参数错误")
	MethodNotAllowed = MakeResult("MethodNotAllowed", "请求方式不允许")
	InternalError    = MakeResult("InternalError", "内部错误")
	ThirdPartError   = MakeResult("ThirdPartError", "第三方错误")
)
