package auth

type LoginByPhoneReq struct {
	Phone   string `json:"phone"`
	SmsCode string `json:"smsCode"`
}

type LoginRes struct {
	UserId string `json:"userId"`
}
