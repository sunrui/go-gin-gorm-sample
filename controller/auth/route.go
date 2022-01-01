package auth

import (
	"medium-server-go/common/gin"
)

func RegisterHandler(server *gin.Starter) {
	server.RegisterHandler("POST", "/login/phone", LoginByPhone)
	server.RegisterHandler("POST", "/login/wechat", LoginByWechat)
}
