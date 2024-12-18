package filter

import (
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/garrettladley/prods/internal/model/category"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"

	go_json "github.com/goccy/go-json"
)

func TestParamsEncode(t *testing.T) {
	tests := []struct {
		name     string
		params   Params
		expected string
	}{
		{
			name:     "empty params",
			params:   Params{},
			expected: "",
		},
		{
			name: "all fields populated",
			params: Params{
				Sort:       Price,
				Order:      Asc,
				Categories: []category.Category{category.Books, category.Electronics},
				Offset:     10,
				Limit:      20,
				PriceMin:   1000,
				PriceMax:   5000,
				StarMin:    300,
				StarMax:    500,
			},
			expected: "categories=books,electronics&limit=20&offset=10&order=asc&price_max=5000&price_min=1000&sort=price&star_max=500&star_min=300",
		},
		{
			name: "only sort and order",
			params: Params{
				Sort:  Name,
				Order: Desc,
			},
			expected: "order=desc&sort=name",
		},
		{
			name: "only categories",
			params: Params{
				Categories: []category.Category{category.Books},
			},
			expected: "categories=books",
		},
		{
			name: "only pagination",
			params: Params{
				Offset: 20,
				Limit:  10,
			},
			expected: "limit=10&offset=20",
		},
		{
			name: "only price range",
			params: Params{
				PriceMin: 1000,
				PriceMax: 2000,
			},
			expected: "price_max=2000&price_min=1000",
		},
		{
			name: "only star range",
			params: Params{
				StarMin: 300,
				StarMax: 400,
			},
			expected: "star_max=400&star_min=300",
		},
		{
			name: "zero values are omitted",
			params: Params{
				Sort:       Price,
				Order:      Asc,
				Categories: []category.Category{},
				Offset:     0,
				Limit:      0,
				PriceMin:   0,
				PriceMax:   0,
				StarMin:    0,
				StarMax:    0,
			},
			expected: "order=asc&sort=price",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.Encode()
			if got != tt.expected {
				t.Errorf("Params.Encode() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestQueryParser(t *testing.T) {
	tests := []struct {
		name           string
		queryParams    string
		expectedParams Params
		expectError    bool
	}{
		{
			name:           "EmptyQuery",
			queryParams:    "",
			expectedParams: Params{},
			expectError:    false,
		},
		{
			name:        "SingleCategory",
			queryParams: "categories=books",
			expectedParams: Params{
				Categories: []category.Category{category.Books},
			},
			expectError: false,
		},
		{
			name:        "MultipleCategories",
			queryParams: "categories=books&categories=electronics",
			expectedParams: Params{
				Categories: []category.Category{
					category.Books,
					category.Electronics,
				},
			},
			expectError: false,
		},
		{
			name:        "AllFieldsValid",
			queryParams: "sort=price&order=asc&categories=books&categories=electronics&offset=10&limit=20&price_min=1000&price_max=5000&star_min=300&star_max=500",
			expectedParams: Params{
				Sort:       "price",
				Order:      "asc",
				Categories: []category.Category{category.Books, category.Electronics},
				Offset:     10,
				Limit:      20,
				PriceMin:   1000,
				PriceMax:   5000,
				StarMin:    300,
				StarMax:    500,
			},
			expectError: false,
		},
		{
			name:           "NumericFieldsWithInvalidValues",
			queryParams:    "offset=abc&limit=def&price_min=xyz&price_max=uvw&star_min=rst&star_max=opq",
			expectedParams: Params{},
			expectError:    true,
		},
	}

	app := fiber.New(fiber.Config{
		JSONEncoder:       go_json.Marshal,
		JSONDecoder:       go_json.Unmarshal,
		ErrorHandler:      xerr.ErrorHandler,
		PassLocalsToViews: true,
	})
	app.Get("/filter", func(c *fiber.Ctx) error {
		params := new(Params)
		if err := c.QueryParser(params); err != nil {
			return xerr.BadRequest(err)
		}
		return c.Status(fiber.StatusOK).JSON(params)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(fiber.MethodGet, "/filter?"+tt.queryParams, nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("failed to make request: %v", err)
			}

			if tt.expectError {
				if resp.StatusCode != fiber.StatusBadRequest {
					t.Errorf("expected status code %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				var p Params
				if err := go_json.NewDecoder(resp.Body).Decode(&p); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				compareParams(t, tt.expectedParams, p)
			}
		})
	}
}

func compareParams(t *testing.T, expected Params, actual Params) {
	t.Helper()

	if expected.Sort != actual.Sort {
		t.Errorf("Sort mismatch: expected %v, got %v", expected.Sort, actual.Sort)
	}
	if expected.Order != actual.Order {
		t.Errorf("Order mismatch: expected %v, got %v", expected.Order, actual.Order)
	}
	if !slices.Equal(expected.Categories, actual.Categories) {
		t.Errorf("Categories mismatch: expected %v, got %v", expected.Categories, actual.Categories)
	}
	if expected.Offset != actual.Offset {
		t.Errorf("Offset mismatch: expected %v, got %v", expected.Offset, actual.Offset)
	}
	if expected.Limit != actual.Limit {
		t.Errorf("Limit mismatch: expected %v, got %v", expected.Limit, actual.Limit)
	}
	if expected.PriceMin != actual.PriceMin {
		t.Errorf("PriceMin mismatch: expected %v, got %v", expected.PriceMin, actual.PriceMin)
	}
	if expected.PriceMax != actual.PriceMax {
		t.Errorf("PriceMax mismatch: expected %v, got %v", expected.PriceMax, actual.PriceMax)
	}
	if expected.StarMin != actual.StarMin {
		t.Errorf("StarMin mismatch: expected %v, got %v", expected.StarMin, actual.StarMin)
	}
	if expected.StarMax != actual.StarMax {
		t.Errorf("StarMax mismatch: expected %v, got %v", expected.StarMax, actual.StarMax)
	}
}
