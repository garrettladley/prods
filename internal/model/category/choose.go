package category

import "github.com/garrettladley/prods/internal/rand"

func ChooseCategories(seed uint64, n uint) []Category {
	return rand.ChooseN(seed, n, Categories[:]...)
}
