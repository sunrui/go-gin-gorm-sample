/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/12 17:42:12
 */

package enum

import (
	"errors"
	"reflect"
)

// 验证码类型
type CodeType string

// 验证码定义
type codeDef struct {
	Login, ResetPassword CodeType
}

// 验证码转换
func (codeDef codeDef) ValueOf(code string) (codeType CodeType, err error) {
	vo := reflect.ValueOf(codeDef)
	typeVo := vo.Type()

	for i := 0; i < vo.NumField(); i++ {
		if typeVo.Field(i).Name == code {
			return vo.Field(i).Interface().(CodeType), nil
		}
	}

	return "", errors.New("code not found")
}

// 验证码实体
var Code = codeDef{
	Login:         "Login",
	ResetPassword: "ResetPassword",
}
