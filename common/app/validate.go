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
	"medium-server-go/common/result"
	"strings"
)

func ValidateParameter(ctx *gin.Context, req interface{}) *result.Result {
	var err error
	if err = ctx.MustBindWith(&req, binding.JSON); err != nil {
		goto haveError
	}

	if err = validator.New().Struct(req); err != nil {
		goto haveError
	}

	return nil

haveError:
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		dataMap := make(map[string]interface{})
		dataMap["error"] = fmt.Sprintf("%s", err)
		return result.ParameterError.WithData(dataMap)
	}

	type ParamError struct {
		Field    string `json:"field"`
		Validate string `json:"validate"`
	}

	var paramErrors []ParamError

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

	return result.ParameterError.WithData(dataMap)
}
