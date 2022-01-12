/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/12 01:07:12
 */

package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"medium-server-go/common/config"
)

// Redis 数据库访问对象
var Redis *redis.Client

// 初始化
func init() {
	redisConfig := config.Get().RedisConfig

	// 数据库连接
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})

	// 使用请求 flush 检测是否连接成功
	err := Redis.FlushDB().Err()
	if err != nil {
		panic(err.Error())
	}
}
