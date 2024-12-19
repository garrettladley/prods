package product

import (
	"testing"
)

func TestProductEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected Product
		actual   Product
		want     bool
	}{
		{
			name:     "same product",
			expected: Products[AirPodsPro.ID],
			actual:   AirPodsPro,
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.expected.Equals(tt.actual); got != tt.want {
				t.Errorf("Product.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
