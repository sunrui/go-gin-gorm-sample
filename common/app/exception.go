/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 07:51:03
 */

package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/config"
	"medium-server-go/common/result"
	"os"
	"runtime"
	"strings"
)

// 异常捕获对象
func exceptionHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 为了更好的调试，在开发环境中输出系统错误。
		if !config.IsDebugMode() {
			// 捕获对象，全部抛出可以使用 panic 方法。
			defer func() {
				if err := recover(); err != nil {
					dataMap := make(map[string]interface{})
					// 判断是否抛出了 result 对象
					res, ok := err.(*result.Result)
					if ok {
						dataMap["error"] = res.Data
					} else {
						dataMap["error"] = err
					}

					mapData := make(map[string]interface{})
					mapData["description"] = err

					type Stack struct {
						Function string
						File     string
					}

					var stacks []Stack

					maxDeep := 6
					pc := make([]uintptr, maxDeep)
					runtime.Callers(4, pc)
					frames := runtime.CallersFrames(pc)

					pwd, _ := os.Getwd()
					pwd = strings.Replace(pwd, "\\", "/", -1)
					goPath := os.Getenv("GOPATH")
					goPath = strings.Replace(goPath, "\\", "/", -1)

					for frame, ok := frames.Next(); ok; frame, ok = frames.Next() {
						file := strings.Replace(frame.File, pwd, "", -1)
						file = strings.Replace(file, goPath, "", -1)
						file = fmt.Sprintf("%s:%d", file, frame.Line)
						function := frame.Function[strings.Index(frame.Function, "/"):]

						stacks = append(stacks, Stack{
							Function: function,
							File:     file,
						})
					}

					dataMap["stacks"] = stacks

					Response(ctx, result.InternalError.WithData(dataMap))
				}
			}()
		}

		ctx.Next()
		handlerFunc(ctx)
	}
}
