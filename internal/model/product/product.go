package product

import (
	"slices"

	"github.com/garrettladley/prods/internal/model/category"
)

type Product struct {
	ID         string              `json:"id"`         // 8 runes, alphanumeric, case sensitive
	Name       string              `json:"name"`       // name of the product
	Categories []category.Category `json:"categories"` // between 1-3 categories
	Stars      uint16              `json:"stars"`      // representation of 0-5 stars, 2 decimal places
	Price      uint32              `json:"price"`      // representation in cents
}

func (p *Product) Equals(other Product) bool {
	if p.ID != other.ID {
		return false
	}

	if p.Name != other.Name {
		return false
	}

	if p.Stars != other.Stars {
		return false
	}

	if p.Price != other.Price {
		return false
	}

	if !slices.Equal(p.Categories, other.Categories) {
		return false
	}

	return true
}
