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

func createMysql(host string, port int, database string, user string, password string) Mysql {
	return Mysql{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	}
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
	Debug = createConfig(
		createMysql("Host", 3316, "medium.db", "User", "Password"),
	)
	Release = createConfig(
		createMysql("Host", 3316, "medium.db", "User", "Password"),
	)
)

func IsDebugMode() bool {
	return os.Getenv("release") == ""
}
