/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/18 19:00:18
 */

package provider

import (
	"github.com/golang-jwt/jwt"
	"medium-server-go/common/config"
)

// 令牌对象
type TokenEntity struct {
	UserId string `json:"userId"`
}

// Jwt 令牌对象
type TokenJwtEntity struct {
	jwt.StandardClaims
	TokenEntity
}

// Jwt 对象定义
type TokenDef struct{}

// Jwt 密钥
var jwtSecret = config.Get().JwtSecret

// 生成 Jwt 字符串
func (tokenDef *TokenDef) Encode(tokenEntity TokenEntity) (string, error) {
	claims := TokenJwtEntity{
		jwt.StandardClaims{},
		tokenEntity,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 验证 Jwt 字符串
func (tokenDef *TokenDef) Decode(token string) (*TokenJwtEntity, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenJwtEntity{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenJwtEntity); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

var Token = &TokenDef{}
