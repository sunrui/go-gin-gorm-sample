/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/12 17:16:12
 */

package service

import (
	"fmt"
	"medium-server-go/controller/sms"
)

type Sms struct{}

func (sms *Sms) Send(phone string, codeType sms.CodeType, sixNumber string) (channel string, reqId string, err error) {
	fmt.Println("send %s, %s, %s", phone, codeType, sixNumber)

	channel = "aliyun"
	reqId = "reqId"
	err = nil
	return
}
