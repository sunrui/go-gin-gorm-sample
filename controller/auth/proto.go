package auth

type LoginByPhoneReq struct {
	Phone   string `json:"phone" binding:"required" validate:"min=11,max=11"`
	SmsCode string `json:"smsCode" binding:"required" validate:"min=6,max=6"`
}

type LoginRes struct {
	UserId string `json:"userId"`
}
