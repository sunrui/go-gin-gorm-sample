/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/01
 */
package app

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/common/config"
	"medium-server-go/common/result"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	engine *gin.Engine
}

func init() {
	if !config.IsDebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}
}

func New() *Server {
	engine := gin.Default()

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, result.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, result.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	engine.Use(jsonResponseMiddleware)
	engine.Use(rateLimitMiddleware(time.Second, 10000, 10000))

	return &Server{
		engine: engine,
	}
}

func (app *Server) Run(port int) {
	err := app.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err.Error())
	}
}
