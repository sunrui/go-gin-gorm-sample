package auth

import (
	"medium-server-go/common/gin"
)

func GetRouter() gin.Router {
	var router gin.Router

	router.GroupName = "/auth"
	router.NeedAuth = true
	router.RouterPaths = []gin.RouterPath{
		{
			HttpMethod:   "POST",
			RelativePath: "/login/phone",
			HandlerFunc:  LoginByPhone,
		}, {
			HttpMethod:   "POST",
			RelativePath: "/login/wechat",
			HandlerFunc:  LoginByWechat,
		},
	}

	return router
}
