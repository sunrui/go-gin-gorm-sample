package auth

import (
	"medium-server-go/common/starter"
)

func RegisterHandler(server *starter.Starter) {
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

	server.RegisterRouter(router)
}
