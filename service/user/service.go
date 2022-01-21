/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/21 00:45:21
 */

package user

import "medium-server-go/framework/db"

// 初始化
func init() {
	// 创建表
	err := db.Mysql.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
}

// 存储数据
func SaveUser(user *User) {
	db.Mysql.Save(user)
}

// 根据 id 获取用户
func FindByPhone(phone string) *User {
	var user User

	query := db.Mysql.First(&user, "phone = ?", phone)
	if query.RowsAffected == 0 {
		return nil
	}

	return &user
}

// 根据 id 获取用户
func FindById(id string) *User {
	var user User

	query := db.Mysql.Find(&user, id)
	if query.RowsAffected == 0 {
		return nil
	}

	return &user
}
