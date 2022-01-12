/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:50:03
 */

package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strings"
)

// 请求参数过滤
func ValidateParameter(ctx *gin.Context, req interface{}) (errData interface{}, ok bool) {
	var validationErrors validator.ValidationErrors
	var err error

	// 默认以 json 方式解析
	if err = ctx.MustBindWith(&req, binding.JSON); err != nil {
		goto ERROR
	}

	// 存在解析参数错误
	if err = validator.New().Struct(req); err != nil {
		goto ERROR
	}

	return nil, true

ERROR:
	// 参数错误对象
	type ParamError struct {
		Field    string `json:"field"`    // 变量名
		Validate string `json:"validate"` // 较验值
	}

	var paramErrors []ParamError

	// 解析内容出错
	if !errors.As(err, &validationErrors) {
		dataMap := make(map[string]interface{})
		dataMap["error"] = fmt.Sprintf("%s", err)

		return dataMap, false
	}

	// 遍历解析参数
	for _, e := range validationErrors {
		validate := e.Tag()
		if len(e.Param()) != 0 {
			validate += "=" + e.Param()
		}

		paramErrors = append(paramErrors, ParamError{
			Field:    strings.ToLower(e.Field()),
			Validate: validate,
		})
	}

	dataMap := make(map[string]interface{})
	dataMap["errors"] = paramErrors

	return dataMap, false
}
