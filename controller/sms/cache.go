/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/13 01:21:13
 */

package sms

import (
	"encoding/json"
	"fmt"
	"medium-server-go/common/db"
	"medium-server-go/enum"
)

// 默认过期时间为 15 分钟
const expiredTime = 15 * 60 * 1000 * 1000

// 缓存数据
type CacheValue struct {
	Code       string `json:"code"`       // 验证码
	VerifyTime int    `json:"verifyTime"` // 验证次数
}

func (cacheValue *CacheValue) ToJson() string {
	marshal, _ := json.Marshal(cacheValue)
	return string(marshal)
}

func (cacheValue *CacheValue) FromJson() {

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
	stringCmd := db.Redis.Get(cache.GetKey())
	if stringCmd.Err() != nil {
		panic(stringCmd.Err().Error())
	}

	cv = CacheValue{}

	return true, CacheValue{}
}

// 增加计数
func (cache *Cache) AddVerifyTime() {

}

// 设置新缓存验证码
func (cache *Cache) Save(code string) {
	cacheValue := CacheValue{
		Code:       code,
		VerifyTime: 0,
	}

	// 设置缓存时间
	statusCmd := db.Redis.Set(cache.GetKey(), cacheValue.ToJson(), expiredTime)
	if statusCmd.Err() != nil {
		panic(statusCmd.Err().Error())
	}

	fmt.Println(statusCmd)
}

// 移除缓存验证码
func (cache *Cache) Remove() {

}
