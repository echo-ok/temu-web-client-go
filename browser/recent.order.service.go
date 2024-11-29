// 最近订单服务
package browser

import (
	"context"

	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/normal"
	"gopkg.in/guregu/null.v4"
)

type recentOrderService struct {
	service
	client *Client
}

type BgOrderQueryParams struct {
	normal.ParameterWithPager
	FulfillmentMode     null.Int    `json:"fulfillmentMode"`
	QueryType           null.Int    `json:"queryType"`
	SortType            null.Int    `json:"sortType"`
	ParentOrderSnList   []string    `json:"parentOrderSnList"`
	TimeZone            null.String `json:"timeZone"`
	ParentAfterSalesTag null.Int    `json:"parentAfterSalesTag"`
	NeedBuySignService  null.Int    `json:"needBuySignService"`
	SellerNoteLabelList []int       `json:"sellerNoteLabelList"`
}

// Query 查询订单列表
func (s recentOrderService) Query(ctx context.Context, params BgOrderQueryParams) (items []entity.RecentOrder, total, totalPages int, isLastPage bool, err error) {
	params.TidyPager()

	var result = struct {
		normal.Response
		Result struct {
			TotalItemNum int                  `json:"totalItemNum"`
			PageItems    []entity.RecentOrder `json:"pageItems"`
		} `json:"result"`
	}{}

	// 设置启用cookie
	s.client.sellerCentralClient.SetCookieJar(s.client.sellerCentralClient.GetClient().Jar)

	resp, err := s.client.sellerCentralClient.R().
		SetResult(&result).
		SetContext(ctx).
		// TODO: 需要从配置中获取
		SetHeader("mallid", "634418212175626").
		SetBody(params).
		Post("/kirogi/bg/mms/recentOrderList")
	if err = recheckError(resp, result.Response, err); err != nil {
		return
	}

	items = result.Result.PageItems
	total, totalPages, isLastPage = parseResponseTotal(params.Page, params.PageSize, result.Result.TotalItemNum)
	if !isLastPage {
		params.Page++
		nextItems, nextTotal, _, _, err := s.Query(ctx, params) // 递归获取
		if err != nil {
			return items, total, totalPages, isLastPage, err
		}
		items = append(items, nextItems...)
		total += nextTotal
	}
	return items, total, totalPages, isLastPage, err
}
