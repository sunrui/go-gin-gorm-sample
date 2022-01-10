/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 03:02:07
 */

package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"medium-server-go/common/config"
	"strings"
	"time"
)

var Default *gorm.DB

func init() {
	sqliteConfig := config.Get().SqliteConfig
	mysqlConfig := config.Get().MysqlConfig
	var err error

	if sqliteConfig != nil {
		Default, err = gorm.Open(sqlite.Open(sqliteConfig.Database), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
	} else if mysqlConfig != nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlConfig.User,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Database)
		Default, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic("no config select")
	}
}

func MakeId() string {
	id := uuid.NewString()
	id = strings.ToUpper(id)
	id = strings.ReplaceAll(id, "-", "")

	return id
}

type Model struct {
	Id        string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
	//CreatedAt time.Time  `json:"created_at"`
	//UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Model) BeforeCreate(tx *gorm.DB) (err error) {
	base.Id = MakeId()

	return nil
}

//
//func (base *Model) BeforeUpdate(tx *gorm.DB) (err error) {
//	base.UpdatedAt = time.Now()
//
//	return nil
//}
