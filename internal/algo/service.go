package algo

import (
	"context"

	"github.com/garrettladley/prods/internal/model/product"
)

type Prompt struct {
	ProductIDs []string `json:"product_ids"`
}

type Solution struct {
	OrderedProductIDs [][]string `json:"ordered_product_ids"`
}

type Service struct {
	numIDs uint
}

func NewService(numIDs uint) *Service {
	return &Service{
		numIDs: numIDs,
	}
}

func (s *Service) Score(ctx context.Context, expected Solution, actual [][]product.Product) int {
	// wrong len, can't even score
	if len(expected.OrderedProductIDs) != len(actual) {
		return -1
	}

	var score int
	for i, expectedIDs := range expected.OrderedProductIDs {
		// just mark as wrong and continue
		if len(expectedIDs) != len(actual[i]) {
			continue
		}

		var correct bool
	inner:
		for exp, expectedID := range expectedIDs {
			expectedProduct := product.Products[expectedID]
			if !expectedProduct.Equals(actual[i][exp]) {
				correct = false
				break inner
			}
		}
		if correct {
			score++
		}
	}

	return score
}

func (s *Service) Solution(ctx context.Context, p Prompt) Solution {
	soln := make([][]string, len(p.ProductIDs))
	pf := NewProductFilter(p.ProductIDs)
	for idx, filter := range Filters {
		soln[idx] = pf.ApplyFilter(filter)
	}
	return Solution{OrderedProductIDs: soln}
}
