package filter

import (
	"net/url"
	"strconv"

	"github.com/garrettladley/prods/internal/model/category"
)

// TODO: rename
type Params struct {
	Sort  SortBy `query:"sort"`
	Order Order  `query:"order"`

	// TODO: will this work with c.QueryParser?
	Categories []category.Category `query:"categories"`

	Offset uint `query:"offset"`
	Limit  uint `query:"limit"`

	PriceMin uint32 `query:"price_min"`
	PriceMax uint32 `query:"price_max"`

	StarMin uint16 `query:"star_min"`
	StarMax uint16 `query:"star_max"`
}

func (p *Params) Encode() string {
	url := make(url.Values)
	if p.Sort != "" {
		url.Add("sort", string(p.Sort))
	}
	if p.Order != "" {
		url.Add("order", string(p.Order))
	}
	if len(p.Categories) > 0 {
		for _, c := range p.Categories {
			url.Add("categories", string(c))
		}
	}
	if p.Offset > 0 {
		url.Add("offset", strconv.Itoa(int(p.Offset)))
	}
	if p.Limit > 0 {
		url.Add("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.PriceMin > 0 {
		url.Add("price_min", strconv.Itoa(int(p.PriceMin)))
	}
	if p.PriceMax > 0 {
		url.Add("price_max", strconv.Itoa(int(p.PriceMax)))
	}
	if p.StarMin > 0 {
		url.Add("star_min", strconv.Itoa(int(p.StarMin)))
	}
	if p.StarMax > 0 {
		url.Add("star_max", strconv.Itoa(int(p.StarMax)))
	}
	return url.Encode()
}
