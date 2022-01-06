/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/04 21:46:04
 */

package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"medium-server-go/common/result"
	"sync"
)

func New(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic(result.InternalError.WithData(err.Error()))
	}

	return db
}

var db *gorm.DB
var dbLock sync.Once

func DB() *gorm.DB {
	if db == nil {
		dbLock.Do(func() {
			db = New("gorm.db")
		})
	}

	return db
}

func Test() {
	type Product struct {
		gorm.Model
		Code  string `gorm:"column:code"`
		Price uint
	}

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	//db.Delete(&product)
}
