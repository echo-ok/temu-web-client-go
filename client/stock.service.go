package client

import (
	"context"
	"fmt"

	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/normal"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type stockService struct {
	service
	client *Client
}

type UpdateMmsBtgProductSalesStockRequestParams struct {
	ProductId          int64            `json:"productId"`
	SkuStockChangeList []SkuStockChange `json:"skuStockChangeList"`
	SkuTypeChangeList  []SkuTypeChange  `json:"skuTypeChangeList,omitempty"`
	IsCheckVersion     bool             `json:"isCheckVersion"`
}

type SkuStockChange struct {
	ProductSkuId          int64  `json:"productSkuId"`
	StockDiff             int64  `json:"stockDiff"`
	CurrentStockAvailable int64  `json:"currentStockAvailable"`
	CurrentShippingMode   int64  `json:"currentShippingMode"`
	WarehouseId           string `json:"warehouseId"`
}

type SkuTypeChange struct {
	ProductSkuId   int64 `json:"productSkuId"`
	ProductSkuType int64 `json:"productSkuType"`
}

type QueryBtgProductStockInfoRequestParams struct {
	ProductId        int64   `json:"productId"`
	ProductSkuIdList []int64 `json:"productSkuIdList"`
}

func (m UpdateMmsBtgProductSalesStockRequestParams) validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ProductId, validation.Required.Error("商品ID不能为空")),
		validation.Field(&m.SkuStockChangeList, validation.Required.Error("SKU库存变更列表不能为空")),
		validation.Field(&m.SkuTypeChangeList, validation.Required.Error("SKU类型变更列表不能为空")),
	)
}

// 查询SKU库存信息 /marvel-mms/cn/api/kiana/starlaod/btg/sales/stock/queryBtgProductStockInfo
func (s stockService) QueryBtgProductStockInfo(ctx context.Context, params QueryBtgProductStockInfoRequestParams) ([]entity.ProductStock, error) {
	var result = struct {
		normal.Response
		Result struct {
			ProductStockList []entity.ProductStock `json:"productStockList"`
		} `json:"result"`
	}{}

	if err := s.client.CheckMallId(); err != nil {
		return nil, err
	}

	resp, err := s.httpClient.R().
		SetHeader("mallid", fmt.Sprintf("%d", s.client.MallId)).
		SetResult(&result).
		SetContext(ctx).
		SetBody(params).
		Post("/marvel-mms/cn/api/kiana/starlaod/btg/sales/stock/queryBtgProductStockInfo")
	if err = recheckError(resp, result.Response, err); err != nil {
		return nil, err
	}

	return result.Result.ProductStockList, nil
}

// 修改库存 /marvel-mms/cn/api/kiana/starlaod/btg/sales/stock/updateMmsBtgProductSalesStock
func (s stockService) UpdateMmsBtgProductSalesStock(ctx context.Context, params UpdateMmsBtgProductSalesStockRequestParams) (bool, error) {
	if err := params.validate(); err != nil {
		return false, err
	}

	var result = struct {
		normal.Response
		Result bool `json:"result"`
	}{}

	resp, err := s.httpClient.R().
		SetResult(&result).
		SetContext(ctx).
		SetBody(params).
		Post("/marvel-mms/cn/api/kiana/starlaod/btg/sales/stock/updateMmsBtgProductSalesStock")
	if err = recheckError(resp, result.Response, err); err != nil {
		return false, err
	}

	return result.Success, nil
}
