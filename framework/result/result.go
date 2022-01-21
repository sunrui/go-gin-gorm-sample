/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package result

import (
	"encoding/json"
)

// 结果对象
type Result struct {
	Code    string      `json:"code"`    // 代码
	Message string      `json:"message"` // 说明
	Data    interface{} `json:"data"`    // 数据
}

// 设置结果对象参数对
func (result Result) WithKeyPair(key string, value string) Result {
	dataMap := make(map[string]string)
	dataMap[key] = value
	result.Data = dataMap
	return result
}

// 设置结果对象数据
func (result Result) WithData(data interface{}) Result {
	result.Data = data
	return result
}

// 重写返回结果对象，使用 json 反序列化
func (result Result) String() string {
	marshal, _ := json.Marshal(result)
	return string(marshal)
}

// 创建结果对象
func CreateResult(code string, message string) Result {
	return Result{
		Code:    code,
		Message: message,
	}
}
