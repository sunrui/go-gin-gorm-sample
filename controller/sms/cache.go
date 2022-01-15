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
type CodeCache struct {
	Code      string `json:"code"`      // 验证码
	ErrVerify int    `json:"errVerify"` // 出错较验次数
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
func (cache *Cache) GetValue() *CodeCache {
	var codeCache CodeCache

	ok := db.Redis.GetJson(cache.GetKey(), &codeCache)
	if ok {
		return &codeCache
	}

	return nil
}

// 设置新缓存验证码
func (cache *Cache) Save(codeCache CodeCache) {
	db.Redis.Set(cache.GetKey(), codeCache, 15*60)
}

// 移除缓存验证码
func (cache *Cache) Del() {
	db.Redis.Del(cache.GetKey())
}
