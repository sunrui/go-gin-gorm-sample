/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/29 18:02:29
 */

package api_area

import (
	"github.com/gin-gonic/gin"
	"medium-server-go/framework/app"
	"medium-server-go/framework/result"
	"medium-server-go/service/biz/area"
	"strconv"
)

// 获取国家
func getCountry(ctx *gin.Context) {
	country := area.GetCountry()
	app.Response(ctx, result.Ok.WithData(country))
}

// 获取省
func getProvince(ctx *gin.Context) {
	provinces := area.GetProvinces()
	app.Response(ctx, result.Ok.WithData(provinces))
}

// 获取市
func getCity(ctx *gin.Context) {
	provinceId, err := strconv.Atoi(ctx.Param("provinceId"))
	if err != nil {
		app.Response(ctx, result.ParameterError.WithKeyPair("provinceId", ctx.Param("provinceId")))
		return
	}

	cities := area.GetCity(provinceId)
	if cities == nil {
		app.Response(ctx, result.NotFound.WithKeyPair("provinceId", provinceId))
		return
	}

	app.Response(ctx, result.Ok.WithData(cities))
}

// 获取地区
func getArea(ctx *gin.Context) {
	cityId, err := strconv.Atoi(ctx.Param("cityId"))
	if err != nil {
		app.Response(ctx, result.ParameterError.WithKeyPair("cityId", ctx.Param("cityId")))
		return
	}

	areas := area.GetArea(cityId)
	if areas == nil {
		app.Response(ctx, result.NotFound.WithKeyPair("cityId", cityId))
		return
	}

	app.Response(ctx, result.Ok.WithData(areas))
}
