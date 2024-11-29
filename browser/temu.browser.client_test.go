package browser

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bestk/temu-helper/config"
	"github.com/bestk/temu-helper/utils"
)

func TestGetAntiContent(t *testing.T) {
	antiContent, err := utils.GetAntiContent()
	if err != nil {
		t.Errorf("获取 Anti-Content 失败: %v", err)
	}

	if antiContent == "" {
		t.Error("获取的 Anti-Content 为空")
	}
	t.Logf("获取的 Anti-Content: %s", antiContent)
}
func TestLogin(t *testing.T) {
	b, err := os.ReadFile("../config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var cfg config.TemuBrowserConfig
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	fmt.Println(cfg)
	temuClient := New(cfg)
	ctx := context.Background()

	loginName := "13825934490"
	password := "Qq123456."

	publicKey, _, err := temuClient.Services.BgAuthService.GetPublicKey()
	if err != nil {
		t.Errorf("获取公钥失败: %v", err)
	}

	encryptPassword, err := utils.EncryptPassword(password, publicKey)
	if err != nil {
		t.Errorf("加密密码失败: %v", err)
	}

	bgLoginRequestParams := BgLoginRequestParams{
		LoginName:       loginName,
		EncryptPassword: encryptPassword,
		KeyVersion:      "1",
	}

	accountId, maskMobile, verifyAuthToken, err := temuClient.Services.BgAuthService.Login(ctx, bgLoginRequestParams)
	if err != nil {
		t.Errorf("登录失败: %v", err)
	}

	if accountId == 0 {
		t.Error("登录成功，但返回的 AccountId 为空")
	}
	t.Logf("登录成功，返回的 AccountId: %d", accountId)
	t.Logf("登录成功，返回的 MaskMobile: %s", maskMobile)
	t.Logf("登录成功，返回的 VerifyAuthToken: %s", verifyAuthToken)
}
