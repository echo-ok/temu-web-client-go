package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bestk/temu-helper/config"
	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/normal"
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

	temuClient := NewClient(cfg)
	ctx := context.Background()

	loginName := ""
	password := ""
	verifyCode := null.NewString("", false)

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
		VerifyCode:      verifyCode.Ptr(),
	}

	accountId, cookies, err := temuClient.Services.BgAuthService.Login(ctx, bgLoginRequestParams)
	if err != nil {
		if err == normal.ErrNeedSMSCode {
			t.Logf("需要短信验证码: %v", err)
			// success, err := temuClient.Services.BgAuthService.GetLoginVerifyCode(ctx, BgGetLoginVerifyCodeRequestParams{
			// 	Mobile: loginName,
			// })
			// if err != nil {
			// 	t.Errorf("获取短信验证码失败: %v", err)
			// }
			// if !success {
			// 	t.Error("获取短信验证码失败")
			// }

			accountId, cookies, err = temuClient.Services.BgAuthService.Login(ctx, bgLoginRequestParams)
			if err != nil {
				t.Errorf("登录失败: %v", err)
				panic(err)
			}
			t.Logf("登录成功，返回的 AccountId: %d", accountId)
			t.Logf("登录成功，返回的 Cookies: %v", cookies)
		} else {
			t.Errorf("登录失败: %v", err)
		}
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
	success, _, err := temuClient.Services.BgAuthService.LoginSellerCentralByCode(ctx, loginByCodeParams)
	if err != nil {
		t.Errorf("登录 Seller Central 失败: %v", err)
		panic(err)
	}
	t.Logf("登录 Seller Central 成功: %v", success)

	userInfo, err := temuClient.Services.BgAuthService.GetSellerCentralUserInfo(ctx)
	if err != nil {
		t.Errorf("获取用户信息失败: %v", err)
		panic(err)
	}
	t.Logf("获取用户信息成功: %+v", userInfo)

	// 设置mallId
	temuClient.SetMallId(634418219352192)

	// 查询订单列表
	params := RecentOrderQueryParams{
		QueryType:           null.NewInt(entity.RecentOrderStatusUnshipped, true),
		FulfillmentMode:     null.NewInt(0, true),
		SortType:            null.NewInt(1, true),
		TimeZone:            null.NewString("UTC+8", true),
		ParentAfterSalesTag: null.NewInt(0, true),
		NeedBuySignService:  null.NewInt(0, true),
		SellerNoteLabelList: []int{},
		ParentOrderSnList:   []string{},
	}
	items, total, _, _, err := temuClient.Services.RecentOrderService.Query(ctx, params)
	if err != nil {
		t.Errorf("查询订单列表失败: %v", err)
	}
	t.Logf("查询订单列表成功，返回的订单数量: %d", total)
	for _, item := range items {
		for _, order := range item.OrderList {
			t.Logf("子订单信息: %+v %+v %+v", order.OrderSn, order.OrderStatus, order.ProductInfoList)
		}
	}
	productId := int64(523914601)
	stockList, err := temuClient.Services.StockService.QueryBtgProductStockInfo(ctx, QueryBtgProductStockInfoRequestParams{
		ProductId:        null.NewInt(productId, true),
		ProductSkuIdList: []int{7634297783},
	})
	if err != nil {
		t.Errorf("查询SKU库存信息失败: %v", err)
		panic(err)
	}
	t.Logf("查询SKU库存信息成功，返回的库存列表: %v", stockList)
	for _, stock := range stockList {
		jsonString, _ := json.Marshal(stock)
		t.Logf("库存信息: %+v", string(jsonString))
	}

}
