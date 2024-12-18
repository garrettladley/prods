package algo

import (
	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/filter"
	"github.com/garrettladley/prods/internal/model/category"
)

func newFilter(opts ...filter.Option) *filter.Params {
	p := filter.New(opts...)
	_ = p.Validate()
	return p
}

var Filters = [...]*filter.Params{
	// basic sort
	newFilter(filter.Sort(filter.Price, filter.Asc)),
	newFilter(filter.Sort(filter.Price, filter.Desc)),
	newFilter(filter.Sort(filter.Name, filter.Asc)),
	newFilter(filter.Sort(filter.Name, filter.Desc)),
	newFilter(filter.Sort(filter.Star, filter.Asc)),
	newFilter(filter.Sort(filter.Star, filter.Desc)),
	// basic pagination
	newFilter(filter.Paginate(0, constants.ProductSubset+5)),
	newFilter(filter.Paginate(1, constants.ProductSubset)),
	newFilter(filter.Paginate(0, constants.ProductSubset/3)),
	newFilter(filter.Paginate(1, constants.ProductSubset/3)),
	newFilter(filter.Paginate(2, constants.ProductSubset/3)),
	// single category
	newFilter(filter.Categories([]category.Category{category.Electronics})),
	newFilter(filter.Categories([]category.Category{category.Apparel})),
	newFilter(filter.Categories([]category.Category{category.HomeGoods})),
	newFilter(filter.Categories([]category.Category{category.Sports})),
	newFilter(filter.Categories([]category.Category{category.Beauty})),
	newFilter(filter.Categories([]category.Category{category.Grocery})),
	newFilter(filter.Categories([]category.Category{category.OfficeSupplies})),
	newFilter(filter.Categories([]category.Category{category.Outdoor})),
	newFilter(filter.Categories([]category.Category{category.Toys})),
	newFilter(filter.Categories([]category.Category{category.Health})),
	newFilter(filter.Categories([]category.Category{category.Automotive})),
	newFilter(filter.Categories([]category.Category{category.Luxury})),
	newFilter(filter.Categories([]category.Category{category.Books})),
	// related category pairs
	newFilter(filter.Categories([]category.Category{category.Electronics, category.OfficeSupplies})),
	newFilter(filter.Categories([]category.Category{category.Beauty, category.Health})),
	newFilter(filter.Categories([]category.Category{category.Sports, category.Outdoor})),
	newFilter(filter.Categories([]category.Category{category.HomeGoods, category.Luxury})),
	// three category combinations
	newFilter(filter.Categories([]category.Category{category.Electronics, category.OfficeSupplies, category.Books})),
	newFilter(filter.Categories([]category.Category{category.Sports, category.Health, category.Beauty})),
	// categories with price buckets
	newFilter(
		filter.Categories([]category.Category{category.Electronics, category.Luxury}),
		filter.PriceBucketer(filter.OverTwentyFive),
	),
	newFilter(
		filter.Categories([]category.Category{category.Grocery, category.HomeGoods}),
		filter.PriceBucketer(filter.UnderFive),
	),
	// categories with star ratings
	newFilter(
		filter.Categories([]category.Category{category.Electronics}),
		filter.StarBucketer(filter.FourPointFivePlus),
	),
	newFilter(
		filter.Categories([]category.Category{category.Beauty, category.Luxury}),
		filter.StarBucketer(filter.FourPlus),
	),
	// complex filter combinations
	newFilter(
		filter.Categories([]category.Category{category.Electronics, category.OfficeSupplies}),
		filter.PriceBucketer(filter.TenToFifteen),
		filter.StarBucketer(filter.FourPlus),
		filter.Sort(filter.Price, filter.Desc),
		filter.Paginate(0, constants.ProductSubset-5),
	),
	newFilter(
		filter.Categories([]category.Category{category.Luxury, category.Beauty, category.Health}),
		filter.PriceBucketer(filter.OverTwentyFive),
		filter.StarBucketer(filter.FourPointFivePlus),
		filter.Sort(filter.Star, filter.Desc),
		filter.Paginate(0, constants.ProductSubset+5),
	),
}

var EncodedFilters = make([]string, len(Filters))

func init() {
	for i, f := range Filters {
		EncodedFilters[i] = f.Encode()
	}
}
