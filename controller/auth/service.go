package auth

import (
	"github.com/gin-gonic/gin"
	starter "medium-server-go/common/gin"
	"net/http"
)

func LoginByPhone(ctx *gin.Context) {
	var req LoginByPhoneReq

	errNo := starter.ValidateParameter(ctx, &req)
	if errNo != nil {
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	ctx.JSON(http.StatusOK, LoginRes{
		UserId: req.Phone,
	})
}

func LoginByWechat(ctx *gin.Context) {

}
