/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/13 01:21:13
 */

package sms

import (
	"fmt"
	"medium-server-go/common/db"
	"medium-server-go/enum"
)

// 缓存数据
type CacheValue struct {
	Code       string `json:"code"`       // 验证码
	VerifyTime int    `json:"verifyTime"` // 验证次数
}

// 缓存对象
type Cache struct {
	Phone    string        `json:"phone"`    // 手机号
	CodeType enum.CodeType `json:"codeType"` // 验证码类型
}

// 获取主键
func (cache *Cache) GetKey() string {
	return fmt.Sprintf("SMS_%s_%s", cache.CodeType, cache.Phone)
}

// 获取缓存的值
func (cache *Cache) GetValue() (ok bool, cv CacheValue) {
	cmd := db.Redis.Get(cache.GetKey())
	if cmd.Err() != nil {
		panic(cmd.Err().Error())
	}

	cv = CacheValue{}

	return true, CacheValue{}
}

// 增加计数
func (cache *Cache) AddVerifyTime() {

}

// 设置新缓存验证码
func (cache *Cache) Save(code string) {
	cmd := db.Redis.Set("hello", "world", 15*60*1000*1000)
	if cmd.Err() != nil {
		panic(cmd.Err().Error())
	}

	fmt.Println(cmd)

	// 设置缓存时间为 15 分钟
	cmd = db.Redis.Set(cache.GetKey(), CacheValue{
		Code:       code,
		VerifyTime: 0,
	}, 15*60*1000*1000)

	if cmd.Err() != nil {
		panic(cmd.Err().Error())
	}

	fmt.Println(cmd)
}

// 移除缓存验证码
func (cache *Cache) Remove() {

}
