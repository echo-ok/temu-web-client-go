package normal

import "errors"

var ErrNeedSMSCode = errors.New("sms code required")

type Response struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
	Result       any    `json:"result"`
}

type ResponseKuajingmaihuo struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
	Result       any    `json:"result"`
}
