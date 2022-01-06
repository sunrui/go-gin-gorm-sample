/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/06 20:20:06
 */

package app

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"medium-server-go/common/result"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("medium.db"), &gorm.Config{})
	if err != nil {
		panic(result.InternalError.WithData(err.Error()))
	}

	Db.Logger.LogMode(logger.Silent)
	fmt.Println("init...database")
}
