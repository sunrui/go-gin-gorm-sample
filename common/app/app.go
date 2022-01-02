/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date:2022/01/01
 */
package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/juju/ratelimit"
	"medium-server-go/common/result"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type App struct {
	engine *gin.Engine
}

func New() *App {
	engine := gin.Default()

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, result.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, result.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	engine.Use(rateLimitMiddleware(time.Second, 10000, 10000))

	return &App{
		engine: engine,
	}
}

type RouterPath struct {
	HttpMethod   string
	RelativePath string
	HandlerFunc  gin.HandlerFunc
}

type Router struct {
	GroupName   string
	NeedAuth    bool
	RouterPaths []RouterPath
}

func (app *App) RegisterRouter(router Router) {
	groupRouter := app.engine.Group(router.GroupName)

	if router.NeedAuth {
		groupRouter.Use(authMiddleware)
	}

	for _, routerPath := range router.RouterPaths {
		switch routerPath.HttpMethod {
		case "GET":
			groupRouter.GET(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "POST":
			groupRouter.POST(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "PUT":
			groupRouter.PUT(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "DELETE":
			groupRouter.DELETE(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		default:
			panic(result.InternalError.WithKeyPair("httpMethod", routerPath.HttpMethod))
		}
	}
}

func (app *App) Run(port int) {
	err := app.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(result.InternalError.WithData(err.Error()))
	}
}

func ValidateParameter(ctx *gin.Context, req interface{}) *result.Result {
	var err error
	if err = ctx.ShouldBind(&req); err != nil {
		goto haveError
	}

	if err = validator.New().Struct(req); err != nil {
		goto haveError
	}

	return nil

haveError:
	type ParamError struct {
		Field    string `json:"field"`
		Validate string `json:"validate"`
	}

	errors := err.(validator.ValidationErrors)
	var paramErrors []ParamError

	for _, e := range errors {
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

func authMiddleware(ctx *gin.Context) {
	ctx.Next()
}

func rateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)

	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			ctx.JSON(http.StatusBadRequest, result.RateLimit)
			return
		}

		ctx.Next()
	}
}

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
