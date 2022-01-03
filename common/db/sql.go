/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/03 23:29:03
 */

package db

import (
	"fmt"
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
	type User struct {
		Id   string
		Name string
		age  int
	}

	var users = []User{{
		Name: "jinzhu1",
		age:  1,
	}, {
		Name: "jinzhu2",
		age:  2,
	}, {
		Name: "jinzhu3",
		age:  3,
	}}

	DB().Create(&users)

	for _, user := range users {
		fmt.Println(user.Id)
	}

}
