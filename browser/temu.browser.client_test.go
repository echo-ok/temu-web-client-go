package browser

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bestk/temu-helper/config"
	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/utils"
	"gopkg.in/guregu/null.v4"
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

// 测试登录
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

	accountId, err := temuClient.Services.BgAuthService.Login(ctx, bgLoginRequestParams)
	if err != nil {
		t.Errorf("登录失败: %v", err)
	}

	if accountId == 0 {
		t.Error("登录成功，但返回的 AccountId 为空")
	}
	t.Logf("登录成功，返回的 AccountId: %d", accountId)

	// 获取验证码
	code, err := temuClient.Services.BgAuthService.ObtainCode(ctx, BgObtainCodeRequestParams{
		RedirectUrl: "https://agentseller.temu.com/main/authentication",
	})
	if err != nil {
		t.Errorf("获取验证码失败: %v", err)
	}
	t.Logf("获取验证码成功，返回的验证码: %s", code)

	loginByCodeParams := BgLoginByCodeRequestParams{
		Code:         code,
		Confirm:      false,
		TargetMallId: 634418212175626,
	}
	success, err := temuClient.Services.BgAuthService.LoginByCode(ctx, loginByCodeParams)
	if err != nil {
		t.Errorf("登录 Seller Central 失败: %v", err)
	}
	t.Logf("登录 Seller Central 成功: %v", success)

	// 查询订单列表
	params := BgOrderQueryParams{
		QueryType:           null.NewInt(entity.RecentOrderStatusUnshipped, true),
		FulfillmentMode:     null.NewInt(0, true),
		SortType:            null.NewInt(1, true),
		TimeZone:            null.NewString("UTC+8", true),
		ParentAfterSalesTag: null.NewInt(0, true),
		NeedBuySignService:  null.NewInt(0, true),
		SellerNoteLabelList: []int{},
		ParentOrderSnList:   []string{},
	}
	items, total, _, _, err := temuClient.Services.BgOrderService.Query(ctx, params)
	if err != nil {
		t.Errorf("查询订单列表失败: %v", err)
	}
	t.Logf("查询订单列表成功，返回的订单数量: %d", total)
	for _, item := range items {
		for _, order := range item.OrderList {
			t.Logf("子订单信息: %+v", order.OrderSn)
		}
	}

}
