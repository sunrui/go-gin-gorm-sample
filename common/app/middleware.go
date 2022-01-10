/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author:sunrui
 * Date: 2022/01/03 07:51:03
 */

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"medium-server-go/common/result"
	"net/http"
	"time"
)

func jsonResponseMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
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
