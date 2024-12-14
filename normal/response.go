package normal

import "errors"

var ErrNeedSMSCode = errors.New("需要短信验证码")
var ErrNeedVerifyCode = errors.New("需要图形验证码")

type Response struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
	Result       any    `json:"result"`
}

type Response2 struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
	Result       any    `json:"result"`
}
