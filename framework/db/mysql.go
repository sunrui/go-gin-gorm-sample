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
	"gorm.io/gorm"
	"medium-server-go/framework/config"
	"strings"
	"time"
)

// Mysql 数据库访问对象
var Mysql *gorm.DB

// 初始化
func init() {
	mysqlConf := config.Get().Mysql
	var err error

	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database)
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

// 创建唯一 id
func CreateUuid() string {
	id := uuid.NewString()
	id = strings.ToUpper(id)
	id = strings.ReplaceAll(id, "-", "")

	return id
}

// 数据库通用对象
type Model struct {
	Id        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"` // 唯一 id
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:milli"`          // 创建时间
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime:milli"`          // 更新时间
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`                          // 删除时间
}

// 创建对象前回调
func (base *Model) BeforeCreate(*gorm.DB) (err error) {
	base.Id = CreateUuid()

	return nil
}
