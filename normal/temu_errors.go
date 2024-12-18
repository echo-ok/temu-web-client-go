package normal

import (
	"errors"
	"fmt"
)

// TemuErrorCode TEMU错误码枚举
type TemuErrorCode struct {
	Code int
	Err  error
}

var (
	TemuErrorNeedSMSCode             TemuErrorCode = TemuErrorCode{Code: 6000001, Err: errors.New("账号发生变更")}   // 需要短信验证码
	TemuErrorNeedVerifyCode          TemuErrorCode = TemuErrorCode{Code: 6000002, Err: errors.New("需要图形验证码")}  // 需要图形验证码
	TemuErrorAccountPasswordNotMatch TemuErrorCode = TemuErrorCode{Code: 40002001, Err: errors.New("账号密码不匹配")} // 账号密码不匹配
)

func (e *TemuErrorCode) Error() error {
	return fmt.Errorf("%d: %s", e.Code, e.Err.Error())
}

var ErrNeedSMSCode = errors.New("需要短信验证码")
var ErrNeedVerifyCode = errors.New("需要图形验证码，请在TEMU后台手动登录，设置Cookie再试")

func GetErrorByCode(code int, message string) TemuErrorCode {
	//  遍历所有错误码，找到对应的错误
	for _, e := range []TemuErrorCode{
		TemuErrorNeedSMSCode,
		TemuErrorNeedVerifyCode,
		TemuErrorAccountPasswordNotMatch,
	} {
		if e.Code == code {
			e.Err = errors.New(message)
			return e
		}
	}

	return TemuErrorCode{
		Code: code,
		Err:  errors.New(message),
	}
}
