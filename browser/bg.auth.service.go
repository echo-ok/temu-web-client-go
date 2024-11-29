// 基础认证服务
package browser

import (
	"context"
	"encoding/json"

	"github.com/bestk/temu-helper/normal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type bgAuthService struct {
	service
	client *Client
}

type BgLoginRequestParams struct {
	LoginName       string `json:"loginName" binding:"required"`
	EncryptPassword string `json:"encryptPassword" binding:"required"`
	KeyVersion      string `json:"keyVersion" default:"1" binding:"required"`
}

type BgObtainCodeRequestParams struct {
	RedirectUrl string `json:"redirectUrl" binding:"required"`
}

type BgLoginByCodeRequestParams struct {
	Code         string `json:"code" binding:"required"`
	Confirm      bool   `json:"confirm" default:"false"`
	TargetMallId uint64 `json:"targetMallId" binding:"required"`
}

func (m BgLoginRequestParams) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.LoginName, validation.Required.Error("登录名不能为空")),
		validation.Field(&m.EncryptPassword, validation.Required.Error("加密密码不能为空")),
	)
}

func (m BgObtainCodeRequestParams) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.RedirectUrl, validation.Required.Error("重定向URL不能为空")),
	)
}

func (m BgLoginByCodeRequestParams) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Code, validation.Required.Error("验证码不能为空")),
	)
}

func (s *bgAuthService) GetPublicKey() (string, string, error) {
	var result = struct {
		normal.Response
		Result struct {
			PublicKey string `json:"publicKey"`
			Version   string `json:"version"`
		} `json:"result"`
	}{}

	resp, err := s.httpClient.R().
		SetResult(&result).
		Post("/bg/quiet/api/mms/key/login")

	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", "", err
	}

	err = recheckError(resp, result.Response, err)
	if err != nil {
		return "", "", err
	}

	return result.Result.PublicKey, result.Result.Version, nil
}

func (s *bgAuthService) Login(ctx context.Context, params BgLoginRequestParams) (uint64, error) {
	if err := params.validate(); err != nil {
		return 0, err
	}

	var result = struct {
		normal.Response
		Result struct {
			MaskMobile      string `json:"maskMobile"`
			VerifyAuthToken string `json:"verifyAuthToken"`
			AccountId       uint64 `json:"accountId"`
		} `json:"result"`
	}{}

	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetBody(params).
		Post("/bg/quiet/api/mms/login")

	if err != nil {
		return 0, err
	}

	err = recheckError(resp, result.Response, err)
	if err != nil {
		return 0, err
	}

	return result.Result.AccountId, nil
}

// ObtainCode 获取验证码 bg/quiet/api/auth/obtainCode
func (s *bgAuthService) ObtainCode(ctx context.Context, params BgObtainCodeRequestParams) (string, error) {
	if err := params.validate(); err != nil {
		return "", err
	}

	var result = struct {
		normal.Response
		Result struct {
			Code string `json:"code"`
		} `json:"result"`
	}{}

	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetBody(params).
		Post("/bg/quiet/api/auth/obtainCode")

	if err != nil {
		return "", err
	}

	err = recheckError(resp, result.Response, err)
	if err != nil {
		return "", err
	}

	return result.Result.Code, nil
}

// LoginSellerCentral 登录 Seller Central
func (s *bgAuthService) LoginSellerCentral(ctx context.Context, url string) (string, error) {
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetDoNotParseResponse(true).
		Get(url)
	if err != nil {
		return "", err
	}
	defer resp.RawResponse.Body.Close()
	return resp.String(), nil
}

// api/seller/auth/loginByCode
func (s *bgAuthService) LoginByCode(ctx context.Context, params BgLoginByCodeRequestParams) (bool, error) {
	if err := params.validate(); err != nil {
		return false, err
	}

	var result = struct {
		normal.Response
		Result struct {
			VerifyAuthToken string `json:"verifyAuthToken"`
		} `json:"result"`
	}{}

	resp, err := s.client.sellerCentralClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetBody(params).
		Post("/api/seller/auth/loginByCode")

	if err != nil {
		return false, err
	}

	err = recheckError(resp, result.Response, err)
	if err != nil {
		return false, err
	}

	return true, nil
}
