package algo

import (
	"context"
	"testing"

	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/model/product"
)

func TestServiceScore(t *testing.T) {
	t.Parallel()

	service := NewService(constants.ProductSubset)

	tests := []struct {
		name     string
		score    int
		expected Solution
		actual   [][]product.Product
	}{
		{
			name:  "sample",
			score: 0,
			expected: Solution{
				OrderedProductIDs: [][]string{
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
			actual: [][]product.Product{
				{product.Products["CLG24680"], product.Products["GAT24680"], product.Products["HPB46802"]},
				{product.Products["CAN24680"], product.Products["PQR35791"], product.Products["DYR13579"]},
				{product.Products["BSE35791"], product.Products["CAN24680"], product.Products["TNT67890"]},
				{product.Products["ABC87429"], product.Products["SNS13579"], product.Products["PQR35791"]},
				{product.Products["CAN24680"], product.Products["VWX46802"], product.Products["STU35791"]},
				{product.Products["DYR13579"], product.Products["ABC87429"], product.Products["BSE35791"]},
				{product.Products["BSE35791"], product.Products["CAN24680"], product.Products["TNT67890"], product.Products["CLG24680"], product.Products["DYR13579"], product.Products["STU35791"], product.Products["GAT24680"], product.Products["PRT46802"], product.Products["HPB46802"], product.Products["LEG24680"], product.Products["HIJ46802"], product.Products["VWX46802"], product.Products["PQR35791"], product.Products["SNS13579"], product.Products["ABC87429"]},
				{product.Products["CAN24680"], product.Products["TNT67890"], product.Products["CLG24680"], product.Products["DYR13579"], product.Products["STU35791"], product.Products["GAT24680"], product.Products["PRT46802"], product.Products["HPB46802"], product.Products["LEG24680"], product.Products["HIJ46802"], product.Products["VWX46802"], product.Products["PQR35791"], product.Products["SNS13579"], product.Products["ABC87429"]},
				{product.Products["BSE35791"], product.Products["CAN24680"], product.Products["TNT67890"], product.Products["CLG24680"], product.Products["DYR13579"]},
				{product.Products["CAN24680"], product.Products["TNT67890"], product.Products["CLG24680"], product.Products["DYR13579"], product.Products["STU35791"]},
				{product.Products["TNT67890"], product.Products["CLG24680"], product.Products["DYR13579"], product.Products["STU35791"], product.Products["GAT24680"]},
				{product.Products["BSE35791"], product.Products["CAN24680"], product.Products["PRT46802"]},
				{product.Products["VWX46802"]},
				{product.Products["STU35791"], product.Products["PQR35791"]},
				{product.Products["GAT24680"], product.Products["HIJ46802"], product.Products["VWX46802"]},
				{product.Products["DYR13579"]},
				{product.Products["CLG24680"], product.Products["GAT24680"]},
				{product.Products["PRT46802"]},
				{product.Products["TNT67890"], product.Products["HIJ46802"]},
				{product.Products["LEG24680"]},
				{product.Products["CLG24680"]},
				{},
				{product.Products["DYR13579"]},
				{product.Products["HPB46802"]},
				{product.Products["BSE35791"], product.Products["CAN24680"], product.Products["PRT46802"]},
				{product.Products["DYR13579"]},
				{product.Products["GAT24680"], product.Products["HIJ46802"], product.Products["VWX46802"]},
				{product.Products["STU35791"], product.Products["PQR35791"]},
				{product.Products["BSE35791"], product.Products["CAN24680"], product.Products["PRT46802"]},
				{product.Products["GAT24680"], product.Products["HIJ46802"], product.Products["VWX46802"]},
				{},
				{product.Products["CLG24680"]},
				{product.Products["BSE35791"], product.Products["SNS13579"], product.Products["ABC87429"]},
				{product.Products["DYR13579"]},
				{},
				{},
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			score := service.Score(context.Background(), tt.expected, tt.actual)
			if score != tt.score {
				t.Errorf("expected %d, got %d", tt.score, score)
			}
		})
	}
}
