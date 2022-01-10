/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/09 14:18:09
 */

package config

import (
	"os"
)

type SqliteConfig struct {
	Database string
}

type MysqlConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type Config struct {
	MysqlConfig  *MysqlConfig
	SqliteConfig *SqliteConfig
}

func createConfig(mysqlConfig *MysqlConfig, sqliteConfig *SqliteConfig) Config {
	return Config{
		MysqlConfig:  mysqlConfig,
		SqliteConfig: sqliteConfig,
	}
}

var (
	debug = createConfig(
		nil,
		&SqliteConfig{
			Database: "medium.db",
		},
	)
	release = createConfig(
		&MysqlConfig{
			Host:     "host",
			Port:     3316,
			Database: "medium.db",
			User:     "user",
			Password: "password",
		},
		nil,
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
