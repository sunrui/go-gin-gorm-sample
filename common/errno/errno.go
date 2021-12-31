package errno

import (
	"encoding/json"
)

type IErrorCode interface {
	WithId(id string) *ErrorCode
	WithData(data interface{}) *ErrorCode
}

type ErrorCode struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (code *ErrorCode) WithKeyPair(key string, value string) *ErrorCode {
	dataMap := make(map[string]string)
	dataMap[key] = value
	code.Data = dataMap
	return code
}

func (code *ErrorCode) WithData(data interface{}) *ErrorCode {
	code.Data = data
	return code
}

func (code *ErrorCode) String() string {
	marshal, err := json.Marshal(code)
	if err != nil {
		return ""
	}

	return string(marshal)
}

func MakeErrorCode(name string, message string) ErrorCode {
	return ErrorCode{
		Code:    name,
		Message: message,
	}
}

var (
	Ok               = MakeErrorCode("Ok", "成功")
	NoAuth           = MakeErrorCode("NoAuth", "没有登录")
	Forbidden        = MakeErrorCode("Forbidden", "没有权限")
	NotFound         = MakeErrorCode("NotFound", "不存在")
	Duplicate        = MakeErrorCode("Duplicate", "已经存在")
	ExceedLimit      = MakeErrorCode("ExceedLimit", "超出限制")
	IllegalData      = MakeErrorCode("IllegalData", "数据不合法")
	RepeatSubmit     = MakeErrorCode("RepeatSubmit", "重复提交")
	ParameterError   = MakeErrorCode("ParameterError", "参数错误")
	MethodNotAllowed = MakeErrorCode("MethodNotAllowed", "请求方式允许")
	InternalError    = MakeErrorCode("InternalError", "内部错误")
	ThirdPartError   = MakeErrorCode("ThirdPartError", "第三方错误")
)
