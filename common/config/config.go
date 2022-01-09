/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"os"
)

type Mysql struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type Config struct {
	Mysql Mysql
}

func createConfig(mysql Mysql) Config {
	return Config{
		Mysql: mysql,
	}
}

var (
	debug = createConfig(
		Mysql{
			Host:     "host",
			Port:     3316,
			Database: "medium.db",
			User:     "user",
			Password: "password",
		},
	)
	release = createConfig(
		Mysql{
			Host:     "host",
			Port:     3316,
			Database: "medium.db",
			User:     "user",
			Password: "password",
		},
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
