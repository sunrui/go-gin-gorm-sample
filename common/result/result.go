/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package result

import "encoding/json"

type IResult interface {
	WithId(id string) *Result
	WithData(data interface{}) *Result
}

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (code *Result) WithKeyPair(key string, value string) *Result {
	dataMap := make(map[string]string)
	dataMap[key] = value
	code.Data = dataMap
	return code
}

func (code *Result) WithData(data interface{}) *Result {
	code.Data = data
	return code
}

func (code *Result) String() string {
	marshal, _ := json.Marshal(code)
	return string(marshal)
}

func MakeResult(code string, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}

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
