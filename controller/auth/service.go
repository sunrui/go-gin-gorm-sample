package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func LoginByPhone(ctx *gin.Context) {

	i := 3
	j := 0
	fmt.Println(i / j)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "login by phone",
	})

	return

	var req LoginByPhoneReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err)
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	// validator.ValidationErrors类型错误则进行翻译
	ctx.JSON(http.StatusOK, gin.H{
		"msg": errs.Error(),
	})

	fmt.Println(req)

	ctx.JSON(http.StatusOK, LoginRes{
		UserId: req.Phone,
	})
}

func LoginByWechat(ctx *gin.Context) {

}
