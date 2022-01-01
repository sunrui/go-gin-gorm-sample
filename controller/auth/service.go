package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/errno"
	"net/http"
)

func LoginByPhone(ctx *gin.Context) {
	i := 3
	j := 0
	k := i / j
	fmt.Println(k)

	ctx.JSON(http.StatusOK, errno.Forbidden)

	//if req.Phone == "15068860507" {
	//	return LoginRes{
	//		UserId: "15068860507",
	//	}
	//}
	//
	//return LoginRes{}
}

func LoginByWechat(ctx *gin.Context) {

}
