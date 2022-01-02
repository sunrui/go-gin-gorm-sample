package auth

import (
	"github.com/gin-gonic/gin"
	starter "medium-server-go/common/gin"
	"medium-server-go/common/result"
	"net/http"
)

func LoginByPhone(ctx *gin.Context) {
	var req LoginByPhoneReq

	errNo := starter.ValidateParameter(ctx, &req)
	if errNo != nil {
		ctx.JSON(http.StatusBadRequest, errNo)
		return
	}

	ctx.JSON(http.StatusOK,
		result.Ok.WithData(LoginRes{
			UserId: req.Phone,
		}))
}

func LoginByWechat(ctx *gin.Context) {

}
