package browser

import (
	"context"

	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/normal"
	"gopkg.in/guregu/null.v4"
)

type bgOrderService struct {
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
func (s bgOrderService) Query(ctx context.Context, params BgOrderQueryParams) (items []entity.BgOrder, total, totalPages int, isLastPage bool, err error) {
	params.TidyPager()

	var result = struct {
		normal.Response
		Result struct {
			Total int              `json:"total"`
			List  []entity.BgOrder `json:"list"`
		} `json:"result"`
	}{}

	// 设置启用cookie
	s.httpClient.SetCookieJar(s.httpClient.GetClient().Jar)

	resp, err := s.httpClient.
		SetBaseURL(s.client.SellerCentralBaseUrl).
		R().
		SetContext(ctx).
		// TODO: 需要从配置中获取
		SetHeader("mallid", "634418212175626").
		SetBody(params).
		SetResult(&result).
		Post("/kirogi/bg/mms/recentOrderList")
	if err = recheckError(resp, result.Response, err); err != nil {
		return
	}

	items = result.Result.List
	total, totalPages, isLastPage = parseResponseTotal(params.Page, params.PageSize, result.Result.Total)

	return
}
