/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package result

import (
	"encoding/json"
)

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
