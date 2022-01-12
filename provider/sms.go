/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/12 17:16:12
 */

package provider

import (
	"fmt"
	"medium-server-go/enum"
)

// 短信对象
type Sms struct{}

// 短信发送
func (sms *Sms) Send(phone string, codeType enum.CodeType, sixNumber string) (channel string, reqId string, err error) {
	fmt.Println("send %s, %s, %s", phone, codeType, sixNumber)

	channel = "aliyun"
	reqId = "reqId"
	err = nil
	return
}
