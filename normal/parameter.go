package normal

import (
	"github.com/bestk/temu-web-client/entity"
)

type Pager struct {
	Page     int `json:"pageNumber"`
	PageSize int `json:"pageSize"`
}

type PagerWithPage struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ParameterWithPager struct {
	Pager
}

type ParameterWithPage struct {
	PagerWithPage
}

// TidyPager 设置翻页数据
func (p *Pager) TidyPager(values ...int) *Pager {
	page := 1
	maxPageSize := entity.MaxPageSize
	n := len(values)
	if n != 0 {
		page = values[0]
		if n >= 2 {
			maxPageSize = values[1]
		}
	}
	if p.Page <= 0 {
		p.Page = page
	}
	if p.PageSize <= 0 || p.PageSize > maxPageSize {
		p.PageSize = maxPageSize
	}
	return p
}

// TidyPager 设置翻页数据
func (p *PagerWithPage) TidyPager(values ...int) *PagerWithPage {
	page := 1
	maxPageSize := entity.MaxPageSize
	n := len(values)
	if n != 0 {
		page = values[0]
		if n >= 2 {
			maxPageSize = values[1]
		}
	}
	if p.Page <= 0 {
		p.Page = page
	}
	if p.PageSize <= 0 || p.PageSize > maxPageSize {
		p.PageSize = maxPageSize
	}
	return p
}
