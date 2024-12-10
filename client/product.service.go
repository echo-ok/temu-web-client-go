package client

import (
	"context"
	"fmt"

	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/normal"
	"gopkg.in/guregu/null.v4"
)

type productService struct {
	service
	client *Client
}

type ProductQueryParams struct {
	normal.ParameterWithPager
	SkcTopStatus null.Int `json:"skcTopStatus"`
}

func (s productService) Query(ctx context.Context, params ProductQueryParams) (items []entity.Product, total, totalPages int, isLastPage bool, err error) {
	params.TidyPager()

	var result = struct {
		normal.Response
		Result struct {
			Total     int              `json:"total"`
			PageItems []entity.Product `json:"pageItems"`
		} `json:"result"`
	}{}

	if err := s.client.CheckMallId(); err != nil {
		return nil, 0, 0, false, err
	}

	resp, err := s.client.BgClient.R().
		SetResult(&result).
		SetContext(ctx).
		SetHeader("mallid", fmt.Sprintf("%d", s.client.MallId)).
		SetBody(params).
		Post("/bg-visage-mms/product/skc/pageQuery")

	if err = recheckError(resp, result.Response, err); err != nil {
		return nil, 0, 0, false, err
	}

	items = result.Result.PageItems
	total, totalPages, isLastPage = parseResponseTotal(params.Page, params.PageSize, result.Result.Total)
	if !isLastPage {
		return nil, 0, 0, false, fmt.Errorf("last page")
	}

	return items, total, totalPages, isLastPage, nil
}
