package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/errno"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusBadRequest, errno.NotFound.WithKeyPair("uri", context.Request.URL.RequestURI()))
	})

	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusBadRequest, errno.MethodNotAllowed.WithKeyPair("uri", context.Request.URL.RequestURI()))
	})

	engine.GET("/hello", response1Handler(HelloHandler))

	err := engine.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// 处理hello请求的handler。如果有异常返回，响应结果也是直接放回
func HelloHandler(ctx *gin.Context) (data interface{}, err error) {
	name := ctx.Query("name")
	if name == "" {
		return nil, errors.New("name is required")
	}

	data = "hello world"
	return
}

type handler func(ctx *gin.Context) (data interface{}, err error)

// 中间件：处理异常和封装响应结果，同时适配gin.HandlerFunc
func response1Handler(h handler) gin.HandlerFunc {
	type response struct {
		Code    int
		Message string
		Data    interface{}
	}
	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			fmt.Println(err)
			ctx.Error(err)

			ctx.JSON(400, "err")

			return
		}
		resp := response{
			Code:    2000000,
			Message: "success",
			Data:    data,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
