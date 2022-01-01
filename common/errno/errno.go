package errno

import (
	"encoding/json"
)

type iErrNo interface {
	WithId(id string) *ErrNo
	WithData(data interface{}) *ErrNo
}

type ErrNo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (code *ErrNo) WithKeyPair(key string, value string) *ErrNo {
	dataMap := make(map[string]string)
	dataMap[key] = value
	code.Data = dataMap
	return code
}

func (code *ErrNo) WithData(data interface{}) *ErrNo {
	code.Data = data
	return code
}

func (code *ErrNo) String() string {
	marshal, _ := json.Marshal(code)
	return string(marshal)
}

func MakeErrNo(name string, message string) ErrNo {
	return ErrNo{
		Code:    name,
		Message: message,
	}
}

var (
	Ok               = MakeErrNo("Ok", "成功")
	NoAuth           = MakeErrNo("NoAuth", "没有登录")
	Forbidden        = MakeErrNo("Forbidden", "没有权限")
	NotFound         = MakeErrNo("NotFound", "不存在")
	Duplicate        = MakeErrNo("Duplicate", "已经存在")
	ExceedLimit      = MakeErrNo("ExceedLimit", "超出限制")
	IllegalData      = MakeErrNo("IllegalData", "数据不合法")
	RepeatSubmit     = MakeErrNo("RepeatSubmit", "重复提交")
	ParameterError   = MakeErrNo("ParameterError", "参数错误")
	MethodNotAllowed = MakeErrNo("MethodNotAllowed", "请求方式允许")
	InternalError    = MakeErrNo("InternalError", "内部错误")
	ThirdPartError   = MakeErrNo("ThirdPartError", "第三方错误")
)
