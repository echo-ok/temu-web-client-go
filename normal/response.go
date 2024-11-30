package normal

import "errors"

var ErrNeedSMSCode = errors.New("sms code required")

type Response struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
	Result       any    `json:"result"`
}
