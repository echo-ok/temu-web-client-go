package browser

import (
	"context"
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type bgAuthService struct {
	service
}

type BgLoginRequestParams struct {
	LoginName       string `json:"loginName" binding:"required"`
	EncryptPassword string `json:"encryptPassword" binding:"required"`
	KeyVersion      string `json:"keyVersion" default:"1" binding:"required"`
}

func (m BgLoginRequestParams) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.LoginName, validation.Required.Error("登录名不能为空")),
		validation.Field(&m.EncryptPassword, validation.Required.Error("加密密码不能为空")),
	)
}

func (s *bgAuthService) GetPublicKey() (string, string, error) {
	var result = struct {
		Response
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

func (s *bgAuthService) Login(ctx context.Context, params BgLoginRequestParams) (uint64, string, string, error) {
	if err := params.validate(); err != nil {
		return 0, "", "", err
	}

	var result = struct {
		Response
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
		return 0, "", "", err
	}

	err = recheckError(resp, result.Response, err)
	if err != nil {
		return 0, "", "", err
	}

	return result.Result.AccountId, result.Result.MaskMobile, result.Result.VerifyAuthToken, nil
}
