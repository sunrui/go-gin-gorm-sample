/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// mysql 配置对象
type mysql struct {
	Host     string // 主机
	Port     int    // 端口
	Database string // 数据库
	User     string // 用户名
	Password string // 密码
}

// redis 配置对象
type redis struct {
	Host     string // 主机
	Port     int    // 端口
	Password string // 密码
	Database int    // 数据库s
}

// 配置对象
type Config struct {
	Mysql        mysql  // Mysql 配置对象
	Redis        redis  // Redis 配置对象
	JwtSecret    []byte // Jwt 密钥
	SmsMagicCode string // 较验短信魔术码
}

// json 反射对象
type jsonConfig struct {
	Environment string
	Debug       Config
	Release     Config
}

// 当前配置
var current *jsonConfig

// 是否在调试环境
func IsDebugMode() bool {
	return current.Environment == "Debug"
}

// 获取当前配置
func Get() *Config {
	if IsDebugMode() {
		return &current.Debug
	} else {
		return &current.Release
	}
}

// 加载当前配置
func init() {
	// 加载配置文件流
	readStream := func() ([]byte, error) {
		// 获取当前项目根目录
		pwd, _ := os.Getwd()
		f, err := os.Open(pwd + "/config.json")
		if err != nil {
			return nil, err
		}

		return ioutil.ReadAll(f)
	}

	stream, err := readStream()
	if err != nil {
		panic(err)
	}

	// 反射配置文件
	err = json.Unmarshal(stream, &current)
	if err != nil {
		panic(err)
	}
}
