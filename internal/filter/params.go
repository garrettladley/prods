package filter

import "github.com/garrettladley/prods/internal/model/category"

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
