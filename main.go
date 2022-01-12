/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2021/12/31
 */
package main

import (
	"medium-server-go/common/app"
	"medium-server-go/common/db"
	"medium-server-go/controller/auth"
	"medium-server-go/controller/sms"
)

func init() {
	statusCmd := db.Redis.Set("hello", "world", 15*60*1000*1000)
	if statusCmd.Err() != nil {
		panic(statusCmd.Err().Error())
	}

	stringCmd := db.Redis.Get("hello")
	if stringCmd.Err() != nil {
		panic(stringCmd.Err().Error())
	}
}

func main() {
	// 创建服务
	server := app.New()

	// 注册路由
	for _, router := range []app.Router{
		sms.GetRouter(),
		auth.GetRouter(),
	} {
		server.RegisterRouter(router)
	}

	// 启动服务
	server.Run(8080)
}
