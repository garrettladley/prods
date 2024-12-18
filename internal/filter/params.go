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

// check if the query params are valid
// for those with default values, set them if they are not provided
func (p *Params) Validate() map[string]string {
	errs := make(map[string]string)

	if p.Sort != "" {
		if !slices.Contains(SortByValues, p.Sort) {
			errs["sort"] = "invalid sort value"
		}
	} else {
		p.Sort = Name
	}

	if p.Order != "" {
		if !slices.Contains(OrderValues, p.Order) {
			errs["order"] = "invalid order value"
		}
	} else {
		p.Order = Asc
	}

	if len(p.Categories) != 0 {
		for _, c := range p.Categories {
			if !slices.Contains(category.Categories[:], c) {
				errs["categories"] = "invalid category value"
				break
			}
		}
	} else {
		p.Categories = category.Categories[:]
	}

	if p.Limit == 0 {
		p.Limit = 3
	}

	if p.PriceMax == 0 {
		p.PriceMax = ^uint32(0)
	}

	if p.PriceMin > p.PriceMax {
		errs["price_min"] = "price_min must be less than price_max"
	}

	if p.StarMax == 0 {
		p.StarMax = 500
	}

	if p.StarMin > p.StarMax {
		errs["star_min"] = "star_min must be less than star_max"
	}

	return errs
}
