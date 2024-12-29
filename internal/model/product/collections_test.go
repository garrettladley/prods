package product

import (
	"testing"

	"github.com/garrettladley/prods/internal/model/category"
)

func TestChooseIDsWithAllCategories(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		seed     uint64
		n        uint
		wantLen  int
		wantNil  bool // should the result be nil
		wantCats bool // should all categories be represented
	}{
		{
			name:     "exact category count",
			seed:     1,
			n:        uint(len(category.Categories)),
			wantLen:  len(category.Categories),
			wantNil:  false,
			wantCats: true,
		},
		{
			name:     "more than categories",
			seed:     2,
			n:        uint(len(category.Categories) + 5),
			wantLen:  len(category.Categories) + 5,
			wantNil:  false,
			wantCats: true,
		},
		{
			name:     "less than categories",
			seed:     3,
			n:        uint(len(category.Categories) - 5),
			wantLen:  0,
			wantNil:  true,
			wantCats: false,
		},
		{
			name:     "zero products",
			seed:     4,
			n:        0,
			wantLen:  0,
			wantNil:  true,
			wantCats: false,
		},
		{
			name:     "large number",
			seed:     5,
			n:        50,
			wantLen:  50,
			wantNil:  false,
			wantCats: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ChooseIDsRepresentingAllCategories(tt.seed, tt.n)

			if len(got) != tt.wantLen {
				t.Errorf("len(ChooseIDsWithAllCategories()) = %v, want %v", len(got), tt.wantLen)
			}

			if tt.wantNil {
				if got != nil {
					t.Errorf("ChooseIDsWithAllCategories() = %v, want nil", got)
				}
			}

			if tt.wantCats {
				catCoverage := make(map[category.Category]bool)
				for _, id := range got {
					prod := Products[id]
					for _, cat := range prod.Categories {
						catCoverage[cat] = true
					}
				}

				if len(catCoverage) != len(category.Categories) {
					t.Errorf("category coverage = %v, want all %v categories covered",
						len(catCoverage), len(category.Categories))
				}
			}

			seen := make(map[string]struct{})
			for _, id := range got {
				if _, ok := seen[id]; ok {
					t.Errorf("ChooseIDsWithAllCategories() = %v, want unique IDs", got)
				}
				seen[id] = struct{}{}
			}
		})
	}
}
