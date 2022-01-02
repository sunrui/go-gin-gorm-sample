/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:51:03
 */

package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/result"
	"net/http"
	"runtime"
	"strings"
)

func catch(ctx *gin.Context) {
	if err := recover(); err != nil {
		funcName, file, line, _ := runtime.Caller(3)
		dataMap := make(map[string]interface{})
		dataMap["error"] = err

		stack := make(map[string]string)
		funcForPCName := runtime.FuncForPC(funcName).Name()
		funcShortName := funcForPCName[strings.LastIndex(funcForPCName, "/")+1:]
		stack["function"] = funcShortName
		file += fmt.Sprintf(":%d", line)
		stack["file"] = file

		dataMap["stack"] = stack
		ret := result.InternalError.WithData(dataMap)

		marshal, _ := json.MarshalIndent(ret, "", "    ")
		fmt.Println(string(marshal))

		file = file[strings.LastIndex(file, "/controller"):]
		stack["file"] = file

		ctx.JSON(http.StatusBadRequest, ret)
	}
}

func catchHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer catch(ctx)
		handlerFunc(ctx)
		ctx.Next()
	}
}