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
	"strings"
)

// mysql 配置对象
type mysql struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Database string `json:"database"` // 数据库
	User     string `json:"user"`     // 用户名
	Password string `json:"password"` // 密码
}

// redis 配置对象
type redis struct {
	Host     string `json:"host"`     // 主机
	Port     int    `json:"port"`     // 端口
	Password string `json:"password"` // 密码
	Database int    `json:"database"` // 数据库s
}

// 配置对象
type Config struct {
	Mysql        mysql  `json:"mysql"`        // Mysql 配置对象
	Redis        redis  `json:"redis"`        // Redis 配置对象
	JwtSecret    []byte `json:"jwtSecret"`    // Jwt 密钥
	SmsMagicCode string `json:"smsMagicCode"` // 较验短信魔术码
}

// json 反射对象
type jsonConfig struct {
	Environment string `json:"environment"` // 当前环境
	Debug       Config `json:"debug"`       // 开发环境
	Release     Config `json:"release"`     //  正式环境
}

// 当前配置
var current *jsonConfig

// 是否在调试环境
func IsDebugMode() bool {
	return strings.ToLower(current.Environment) == "debug"
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
