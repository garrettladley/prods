package algo

import (
	"testing"

	"github.com/garrettladley/prods/internal/filter"
	"github.com/garrettladley/prods/internal/model/product"
)

func TestApplyFilter(t *testing.T) {
	t.Parallel()

	// TODO: add more
	tests := []struct {
		name   string
		ids    []string
		params *filter.Params
		want   []string
	}{
		{
			name:   "no filter",
			ids:    ids[:],
			params: newFilter(),
			want:   ids[:],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewProductFilter(tt.ids).ApplyFilter(tt.params)

			if len(actual) != len(tt.want) {
				t.Fatalf("len(actual) = %d, len(tt.want) = %d", len(actual), len(tt.want))
			}

			for i := range actual {
				if actual[i] != tt.want[i] {
					t.Errorf("actual[%d] = %s, tt.want[%d] = %s", i, actual[i], i, tt.want[i])
				}
			}
		})
	}
}

var ids = [...]string{
	product.AirPodsPro.ID,            // Electronics
	product.LeviBluejeans.ID,         // Apparel
	product.TreadmillMachine.ID,      // Sports, HomeGoods
	product.ToiletPaper.ID,           // Grocery, HomeGoods
	product.LogitechMouse.ID,         // Electronics, OfficeSupplies
	product.SurfboardBrand.ID,        // Sports, Outdoor
	product.LegoCreatorExpert.ID,     // Toys
	product.TheragunElite.ID,         // Health, Sports
	product.TomFordCologne.ID,        // Beauty
	product.GoodyearTires.ID,         // Automotive
	product.RolexSubmariner.ID,       // Luxury
	product.HarryPotterBoxSet.ID,     // Books
	product.PamperedChefStoneware.ID, // HomeGoods
	product.BanwoodBalanceBike.ID,    // Outdoor, Toys
	product.ColgateToothpaste.ID,     // Health, Grocery
}
