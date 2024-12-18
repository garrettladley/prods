package filter

import (
	"net/url"
	"slices"
	"strconv"

	"github.com/garrettladley/prods/internal/model/category"
)

// TODO: rename
type Params struct {
	Sort  SortBy `query:"sort"`
	Order Order  `query:"order"`

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

func (p *Params) Validate() map[string]string {
	errs := make(map[string]string)

	if !slices.Contains(SortByValues, p.Sort) {
		errs["sort"] = "invalid sort value"
	}

	if !slices.Contains(OrderValues, p.Order) {
		errs["order"] = "invalid order value"
	}

	if len(p.Categories) > 0 {
		for _, c := range p.Categories {
			if !slices.Contains(category.Categories[:], c) {
				errs["categories"] = "invalid category value"
				break
			}
		}
	}

	if p.PriceMin > p.PriceMax {
		errs["price_min"] = "price_min must be less than price_max"
	}

	if p.StarMin > p.StarMax {
		errs["star_min"] = "star_min must be less than star_max"
	}

	return errs
}
