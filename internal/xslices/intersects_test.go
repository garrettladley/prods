package xslices

import "testing"

func TestIntersects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		base      []string
		candidate []string
		want      bool
	}{
		{
			name:      "empty slices",
			base:      []string{},
			candidate: []string{},
			want:      false,
		},
		{
			name:      "empty base",
			base:      []string{},
			candidate: []string{"a", "b"},
			want:      false,
		},
		{
			name:      "empty candidate",
			base:      []string{"a", "b"},
			candidate: []string{},
			want:      false,
		},
		{
			name:      "no overlap",
			base:      []string{"a", "b"},
			candidate: []string{"c", "d"},
			want:      false,
		},
		{
			name:      "single element overlap",
			base:      []string{"a", "b"},
			candidate: []string{"b", "c"},
			want:      true,
		},
		{
			name:      "multiple element overlap",
			base:      []string{"a", "b", "c"},
			candidate: []string{"b", "c", "d"},
			want:      true,
		},
		{
			name:      "identical slices",
			base:      []string{"a", "b"},
			candidate: []string{"a", "b"},
			want:      true,
		},
		{
			name:      "single element slices with overlap",
			base:      []string{"a"},
			candidate: []string{"a"},
			want:      true,
		},
		{
			name:      "single element slices without overlap",
			base:      []string{"a"},
			candidate: []string{"b"},
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Intersects(tt.base, tt.candidate)
			if got != tt.want {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}
