/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"os"
)

type MysqlConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type Config struct {
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
}

func createConfig(mysqlConfig MysqlConfig, redisConfig RedisConfig) Config {
	return Config{
		MysqlConfig: mysqlConfig,
		RedisConfig: redisConfig,
	}
}

var (
	debug = createConfig(
		MysqlConfig{
			Host:     "localhost",
			Port:     3306,
			Database: "medium",
			User:     "root",
			Password: "root",
		},
		RedisConfig{},
	)
	release = createConfig(
		MysqlConfig{
			Host:     "localhost",
			Port:     3306,
			Database: "medium",
			User:     "root",
			Password: "root",
		},
		RedisConfig{},
	)
)

func Get() Config {
	if IsDebugMode() {
		return debug
	} else {
		return release
	}
}

func IsDebugMode() bool {
	return os.Getenv("release") == ""
}
