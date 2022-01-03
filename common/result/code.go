/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 19:03:03
 */

package result

var (
	Ok               = makeResult("Ok", "成功")
	NoAuth           = makeResult("NoAuth", "没有登录")
	Duplicate        = makeResult("Duplicate", "已经存在")
	Forbidden        = makeResult("Forbidden", "没有权限")
	NotFound         = makeResult("NotFound", "不存在")
	NotMatch         = makeResult("NotMatch", "不匹配")
	RateLimit        = makeResult("ExceedLimit", "超出限制")
	LogicError       = makeResult("LogicError", "逻辑错误")
	ParameterError   = makeResult("ParameterError", "参数错误")
	MethodNotAllowed = makeResult("MethodNotAllowed", "请求方式不允许")
	InternalError    = makeResult("InternalError", "内部错误")
	ThirdPartError   = makeResult("ThirdPartError", "第三方错误")
)
