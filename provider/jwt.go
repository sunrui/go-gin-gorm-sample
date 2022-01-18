/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 17:26:03
 */

package provider

import (
	jwt "github.com/dgrijalva/jwt-go"
	"medium-server-go/common/config"
	"time"
)

// Jwt 对象
type JwtEntity struct {
	jwt.StandardClaims
	userId string
}

// Jwt 对象定义
type JwtDef struct{}

// Jwt 密钥
var jwtSecret = config.Get().JwtSecret

// 生成 Jwt 字符串
func (jwtDef *JwtDef) Encode(userId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := JwtEntity{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
		userId,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证 Jwt 字符串
func Decode(token string) (*JwtEntity, error) {
	jwtClaims, err := jwt.ParseWithClaims(token, &JwtEntity{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if jwtClaims != nil {
		if claims, ok := jwtClaims.Claims.(*JwtEntity); ok && jwtClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

var Jwt = &JwtDef{}
