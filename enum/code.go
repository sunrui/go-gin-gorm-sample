/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/12 17:42:12
 */

package enum

import "reflect"

// 验证码类型
type CodeType string

// 验证码定义
type codeDef struct {
	Login, ResetPassword CodeType
}

// 验证码转换
func (codeDef codeDef) ValueOf(id string) (codeType CodeType, ok bool) {
	vo := reflect.ValueOf(codeDef)
	typeVo := vo.Type()

	for i := 0; i < vo.NumField(); i++ {
		if typeVo.Field(i).Name == id {
			return vo.Field(i).Interface().(CodeType), true
		}
	}

	return "", false
}

// 验证码实体
var Code = codeDef{
	Login:         "Login",
	ResetPassword: "ResetPassword",
}
