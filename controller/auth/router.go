package auth

import (
	"medium-server-go/common/starter"
)

func GetRouter() starter.Router {
	var router starter.Router

	router.GroupName = "/auth"
	router.NeedAuth = true
	router.RouterPaths = []starter.RouterPath{
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
