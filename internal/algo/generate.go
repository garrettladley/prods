package algo

import (
	"github.com/garrettladley/prods/internal/model/product"
)

func (s *Service) Generate(seed uint64) *Prompt {
	return &Prompt{
		ProductIDs: product.ChooseIDs(seed, s.numIDs),
	}
}
