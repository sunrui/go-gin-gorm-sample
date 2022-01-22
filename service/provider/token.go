/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/18 19:00:18
 */

package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"medium-server-go/framework/config"
	"strings"
)

// 令牌对象
type TokenEntity struct {
	UserId string `json:"userId"`
}

// Jwt 令牌对象
type tokenJwtEntity struct {
	jwt.StandardClaims
	TokenEntity
}

// Jwt 对象定义
type tokenDef struct{}

// jwt 密钥
var jwtSecret = config.Get().JwtSecret

// 令牌 key 名称
const tokenKey = "token"

// 生成 Jwt 字符串
func encode(tokenEntity TokenEntity) (token string, err error) {
	claims := tokenJwtEntity{
		jwt.StandardClaims{},
		tokenEntity,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 验证 Jwt 字符串
func decode(token string) (tokenEntity *TokenEntity, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &tokenJwtEntity{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*tokenJwtEntity); ok && tokenClaims.Valid {
			return &claims.TokenEntity, nil
		}
	}

	return nil, err
}

// 写入 cookie 令牌
func (*tokenDef) WriteToken(ctx *gin.Context, userId string, maxAge int) {
	// 生成用户令牌
	token, err := encode(TokenEntity{
		UserId: userId,
	})
	if err != nil {
		return
	}

	// 写入令牌，默认 30 天
	ctx.SetCookie(tokenKey, token, maxAge,
		"/", "localhost", false, true)
}

// 获取当前用户令牌
func (*tokenDef) GetTokenEntity(ctx *gin.Context) (tokenEntity *TokenEntity, err error) {
	var token string

	// 从 header 中获取令牌
	getHeaderToken := func() string {
		token = ctx.GetHeader("Authorization")
		if token == "" {
			return ""
		}

		prefix := "Bearer "
		if strings.Index(token, prefix) != 0 {
			return ""
		}

		return token[len(prefix):]
	}

	token = getHeaderToken()
	if token == "" {
		// 从 cookie 中取出 token
		token, err = ctx.Cookie(tokenKey)
		if err != nil {
			return nil, err
		}
	}

	return decode(token)
}

// 移除令牌
func (*tokenDef) RemoveToken(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
}

var Token = &tokenDef{}
