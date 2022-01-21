/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 19:03:03
 */

package result

// 通用返回对象码
var (
	Ok               = CreateResult("Ok", "成功")
	NoAuth           = CreateResult("NoAuth", "没有登录")
	Duplicate        = CreateResult("Duplicate", "已经存在")
	Forbidden        = CreateResult("Forbidden", "没有权限")
	NotFound         = CreateResult("NotFound", "不存在")
	NotMatch         = CreateResult("NotMatch", "不匹配")
	RateLimit        = CreateResult("ExceedLimit", "超出限制")
	LogicError       = CreateResult("LogicError", "逻辑错误")
	ParameterError   = CreateResult("ParameterError", "参数错误")
	MethodNotAllowed = CreateResult("MethodNotAllowed", "请求方式不允许")
	InternalError    = CreateResult("InternalError", "内部错误")
	ThirdPartError   = CreateResult("ThirdPartError", "第三方错误")
)
