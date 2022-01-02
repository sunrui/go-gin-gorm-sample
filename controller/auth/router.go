package auth

import (
	"medium-server-go/common/gin"
)

func GetRouter() app.Router {
	var router app.Router

	router.GroupName = "/auth"
	router.NeedAuth = true
	router.RouterPaths = []app.RouterPath{
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
