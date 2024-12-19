package algo

import (
	"testing"

	"github.com/garrettladley/prods/internal/filter"
	"github.com/garrettladley/prods/internal/model/product"
)

func TestApplyFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		ids    []string
		params []*filter.Params
		want   [][]string
	}{
		{
			name:   "sample",
			ids:    []string{"HIJ46802", "DYR13579", "PRT46802", "GAT24680", "LEG24680", "CLG24680", "VWX46802", "SNS13579", "HPB46802", "PQR35791", "TNT67890", "ABC87429", "STU35791", "CAN24680", "BSE35791"},
			params: Filters[:],
			want: [][]string{
				{"CLG24680", "GAT24680", "HPB46802"},
				{"CAN24680", "PQR35791", "DYR13579"},
				{"BSE35791", "CAN24680", "TNT67890"},
				{"ABC87429", "SNS13579", "PQR35791"},
				{"CAN24680", "VWX46802", "STU35791"},
				{"DYR13579", "ABC87429", "BSE35791"},
				{"BSE35791", "CAN24680", "TNT67890", "CLG24680", "DYR13579", "STU35791", "GAT24680", "PRT46802", "HPB46802", "LEG24680", "HIJ46802", "VWX46802", "PQR35791", "SNS13579", "ABC87429"},
				{"CAN24680", "TNT67890", "CLG24680", "DYR13579", "STU35791", "GAT24680", "PRT46802", "HPB46802", "LEG24680", "HIJ46802", "VWX46802", "PQR35791", "SNS13579", "ABC87429"},
				{"BSE35791", "CAN24680", "TNT67890", "CLG24680", "DYR13579"},
				{"CAN24680", "TNT67890", "CLG24680", "DYR13579", "STU35791"},
				{"TNT67890", "CLG24680", "DYR13579", "STU35791", "GAT24680"},
				{"BSE35791", "CAN24680", "PRT46802"},
				{"VWX46802"},
				{"STU35791", "PQR35791"},
				{"GAT24680", "HIJ46802", "VWX46802"},
				{"DYR13579"},
				{"CLG24680", "GAT24680"},
				{"PRT46802"},
				{"TNT67890", "HIJ46802"},
				{"LEG24680"},
				{"CLG24680"},
				{},
				{"DYR13579"},
				{"HPB46802"},
				{"BSE35791", "CAN24680", "PRT46802"},
				{"DYR13579"},
				{"GAT24680", "HIJ46802", "VWX46802"},
				{"STU35791", "PQR35791"},
				{"BSE35791", "CAN24680", "PRT46802"},
				{"GAT24680", "HIJ46802", "VWX46802"},
				{},
				{"CLG24680"},
				{"BSE35791", "SNS13579", "ABC87429"},
				{"DYR13579"},
				{},
				{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for i, p := range tt.params {
				got := NewProductFilter(tt.ids).ApplyFilter(p)
				if len(got) != len(tt.want[i]) {
					t.Errorf("ApplyFilter() = %v, want %v", got, tt.want[i])
				}
				for j, id := range got {
					if id != tt.want[i][j] {
						t.Errorf("ApplyFilter() = %v, want %v", got, tt.want[i])
					}
				}
			}
		})
	}
}

//nolint:unused
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
