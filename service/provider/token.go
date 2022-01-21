/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/18 19:00:18
 */

package provider

import (
	"github.com/golang-jwt/jwt"
	"medium-server-go/framework/config"
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
var jwtSecret = config.Conf.Config().JwtSecret

// 生成 Jwt 字符串
func (*tokenDef) Encode(tokenEntity TokenEntity) (token string, err error) {
	claims := tokenJwtEntity{
		jwt.StandardClaims{},
		tokenEntity,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 验证 Jwt 字符串
func (*tokenDef) Decode(token string) (tokenEntity *TokenEntity, err error) {
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

var Token = &tokenDef{}
