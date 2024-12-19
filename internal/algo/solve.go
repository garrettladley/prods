package algo

import (
	"sort"

	"github.com/garrettladley/prods/internal/filter"
	"github.com/garrettladley/prods/internal/model/product"
	"github.com/garrettladley/prods/internal/xslices"
)

type ProductFilter struct {
	products []product.Product
}

var AllProducts = NewProductFilter(product.IDs[:])

func NewProductFilter(ids []string) *ProductFilter {
	products := make([]product.Product, len(ids))
	for i, id := range ids {
		products[i] = product.Products[id]
	}
	return &ProductFilter{products: products}
}

func (pf *ProductFilter) ApplyFilter(p *filter.Params) []string {
	productsCopy := make([]product.Product, len(pf.products))
	copy(productsCopy, pf.products)

	filtered := pf.filterProducts(p)

	if p.Sort != "" {
		pf.sortProducts(filtered, p.Sort, p.Order)
	}

	return pf.applyPagination(filtered, p.Offset, p.Limit)
}

func (pf ProductFilter) Solve(params []*filter.Params) [][]string {
	results := make([][]string, len(params))
	for i, p := range params {
		results[i] = pf.ApplyFilter(p)
	}
	return results
}

func (pf *ProductFilter) filterProducts(p *filter.Params) []product.Product {
	filtered := make([]product.Product, 0, len(pf.products))
	for _, prod := range pf.products {
		if pf.meetsFilterCriteria(&prod, p) {
			filtered = append(filtered, prod)
		}
	}
	return filtered
}

func (pf *ProductFilter) meetsFilterCriteria(prod *product.Product, p *filter.Params) bool {
	if prod.Price < p.PriceMin || prod.Price > p.PriceMax {
		return false
	}
	if prod.Stars < p.StarMin || prod.Stars > p.StarMax {
		return false
	}
	if len(p.Categories) > 0 && !xslices.Intersects(prod.Categories, p.Categories) {
		return false
	}
	return true
}

func (pf *ProductFilter) sortProducts(products []product.Product, sortBy filter.SortBy, order filter.Order) {
	sort.Slice(products, func(i, j int) bool {
		less := pf.compareProducts(products[i], products[j], sortBy)
		if order == filter.Desc {
			return !less
		}
		return less
	})
}

func (pf *ProductFilter) compareProducts(a product.Product, b product.Product, sortBy filter.SortBy) bool {
	switch sortBy {
	case filter.Price:
		return a.Price < b.Price
	case filter.Name:
		return a.Name < b.Name
	case filter.Stars:
		return a.Stars < b.Stars
	default:
		return true // default to original order
	}
}

func (pf *ProductFilter) applyPagination(products []product.Product, offset uint, limit uint) []string {
	start := int(offset)
	end := start + int(limit)

	if start > len(products) {
		start = len(products)
	}
	if end > len(products) {
		end = len(products)
	}

	results := make([]string, end-start)
	for i, prod := range products[start:end] {
		results[i] = prod.ID
	}
	return results
}
