package product

import "github.com/garrettladley/prods/internal/model/category"

type Product struct {
	ID         string              // 8 runes, alphanumeric, case sensitive
	Name       string              // name of the product
	Categories []category.Category // between 1-3 categories
	Stars      uint16              // representation of 0-5 stars, 2 decimal places
	Price      uint32              // representation in cents
}
