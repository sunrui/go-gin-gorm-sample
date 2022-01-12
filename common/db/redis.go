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

var Redis *redis.Client

func init() {
	redisConfig := config.Get().RedisConfig

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})

	err := Redis.FlushDB().Err()
	if err != nil {
		panic(err.Error())
	}
}
