/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"os"
)

// Mysql 配置对象
type mysql struct {
	Host     string // 主机
	Port     int    // 端口
	Database string // 数据库
	User     string // 用户名
	Password string // 密码
}

// Redis 配置对象
type redis struct {
	Host     string // 主机
	Port     int    // 端口
	Password string // 密码
	Database int    // 数据库s
}

// 配置对象
type Config struct {
	Mysql     mysql  // Mysql 配置对象
	Redis     redis  // Redis 配置对象
	JwtSecret []byte // Jwt 密钥
}

var (
	// 调试环境配置
	debug = Config{
		Mysql: mysql{
			Host:     "localhost",
			Port:     3306,
			Database: "medium",
			User:     "root",
			Password: "root",
		},
		Redis: redis{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Database: 0,
		},
		JwtSecret: []byte("39184F28-F1D8-46CE-A3EB-7331F0D299B0"),
	}

	// 正式环境配置
	release = Config{
		Mysql: mysql{
			Host:     "localhost",
			Port:     3306,
			Database: "medium",
			User:     "root",
			Password: "root",
		},
		Redis: redis{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Database: 0,
		},
		JwtSecret: []byte("39184F28-F1D8-46CE-A3EB-7331F0D299B0"),
	}
)

// 获取当前配置
func Get() *Config {
	if IsDebugMode() {
		return &debug
	} else {
		return &release
	}
}

// 是否调试模式
func IsDebugMode() bool {
	return os.Getenv("release") == ""
}
