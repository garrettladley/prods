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

	var numCorrect int
	for i, expectedIDs := range expected.OrderedProductIDs {
		// just mark as wrong and continue
		if len(expectedIDs) != len(actual[i]) {
			continue
		}

		correct := true
	inner:
		for exp, expectedID := range expectedIDs {
			expectedProduct := product.Products[expectedID]
			if !expectedProduct.Equals(actual[i][exp]) {
				correct = false
				break inner
			}
		}
		if correct {
			numCorrect++
		}
	}

	return len(expected.OrderedProductIDs) - numCorrect
}

func (s *Service) Solution(ctx context.Context, p Prompt) Solution {
	pf := NewProductFilter(p.ProductIDs)
	return Solution{OrderedProductIDs: pf.Solve(Filters[:])}
}
