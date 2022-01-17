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
type codeCache struct {
	Code      string `json:"code"`      // 验证码
	ErrVerify int    `json:"errVerify"` // 出错较验次数
}

// 缓存对象
type Cache struct {
	Phone    string        `json:"phone"`    // 手机号
	CodeType enum.CodeType `json:"codeType"` // 验证码类型
}

// 获取主键
func (cache *Cache) getKey() string {
	return fmt.Sprintf("SMS_%s_%s", cache.CodeType, cache.Phone)
}

// 获取缓存的值
func (cache *Cache) getValue() *codeCache {
	var codeCache codeCache

	ok := db.Redis.GetJson(cache.getKey(), &codeCache)
	if ok {
		return &codeCache
	}

	return nil
}

// 获取缓存是否存在
func (cache *Cache) Exists() bool {
	return db.Redis.Exists(cache.getKey())
}

// 设置新缓存验证码
func (cache *Cache) save(codeCache codeCache) {
	db.Redis.Set(cache.getKey(), codeCache, 15*60)
}

// 移除缓存验证码
func (cache *Cache) Del() {
	db.Redis.Del(cache.getKey())
}

// 较验验证码
func (cache *Cache) Verify(code string) bool {
	// 获取缓存数据
	value := cache.getValue()
	if value == nil {
		return false
	}

	// 如果验证码较验错误
	if value.Code != code {
		// 增加缓存引用记数
		value.ErrVerify += 1

		// 如果已经较验出错 5 次，移除现有验证码
		if value.ErrVerify == 5 {
			cache.Del()
		} else {
			// 更新出错较验次数
			cache.save(*value)
		}

		return false
	}

	return true
}
