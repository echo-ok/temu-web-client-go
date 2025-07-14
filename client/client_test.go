package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/bestK/temu-web-client/config"
	"github.com/bestK/temu-web-client/entity"
	"github.com/bestK/temu-web-client/normal"
	"github.com/bestK/temu-web-client/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

var temuClient *Client

var ctx = context.Background()

func TestMain(m *testing.M) {
	b, err := os.ReadFile("../config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var cfg config.TemuBrowserConfig
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	temuClient = NewClient(cfg)
	m.Run()
}

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
	loginName := ""
	password := ""
	mallId := 0
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
		if errors.Is(err, normal.ErrNeedSMSCode) {
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
		TargetMallId: mallId,
	}
	temuClient.SetMallId(mallId)
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
}

func TestRecentOrder(t *testing.T) {
	TestLogin(t)
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

func TestCustomizedInformation(t *testing.T) {
	TestLogin(t)
	gotItems, gotTotal, gotTotalPages, gotIsLastPage, err := temuClient.Services.CustomizedInformationService.Query(ctx, CustomizedInformationQueryParams{SubPurchaseOrderSns: []string{"WB2507101860720"}})
	assert.Equal(t, err, nil)
	assert.Equal(t, 2, len(gotItems))
	assert.Equal(t, 2, gotTotal)
	assert.Equal(t, 1, gotTotalPages)
	assert.Equal(t, true, gotIsLastPage)
}

func TestFinanceAccountFunds(t *testing.T) {
	TestLogin(t)
	accountFunds, err := temuClient.Services.FinanceService.AccountFunds(ctx)
	assert.Equal(t, err, nil)
	assert.Equal(t, "13950.20", accountFunds.TotalAmount)
}
