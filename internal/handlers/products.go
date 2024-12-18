package handlers

import (
	"errors"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/filter"
	"github.com/garrettladley/prods/internal/model/product"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
)

// Products godoc
//
//	@Summary		Retrieve filtered products
//	@Description	Fetches a list of products based on various filter parameters such as sorting, categories, price range, and star ratings.
//	@Description	By default, all categories are included if none are specified.
//	@Description	For categories, replace spaces with %20 and join with commas. For example, [office supplies, electronics ] becomes "?categories=office%20supplies,electronics".
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			sort		query		string			false	"Sort by field"												Enums(name, price, stars)	default(name)
//	@Param			order		query		string			false	"Sort order"												Enums(asc, desc)			default(asc)
//	@Param			categories	query		[]string		false	"Filter by categories"										Enums(electronics, apparel, home goods, sports, beauty, grocery, office supplies, outdoor, toys, health, automotive, luxury, books)"
//	@Param			offset		query		uint			false	"Number of pages of size limit to skip"									default(0)
//	@Param			limit		query		uint			false	"Maximum number of items to return"							default(3)
//	@Param			price_min	query		uint32			false	"Minimum price in cents"									default(0)
//	@Param			price_max	query		uint32			false	"Maximum price in cents"									default(4294967295)
//	@Param			star_min	query		uint16			false	"Minimum star rating (e.g., 400 for 4.00 stars)"			default(0)
//	@Param			star_max	query		uint16			false	"Maximum star rating (e.g., 500 for 5.00 stars)"			default(500)
//	@Success		200			{array}		product.Product	"Filtered list of products"
//	@Failure		400			{object}	xerr.APIError	"Invalid query parameters or request data"
//	@Failure		429			{object}	xerr.APIError	"Too many requests"
//	@Failure		500			{object}	xerr.APIError	"Internal server error"
//	@Router			/api/v1/products [get]
func (s *Service) Products(c *fiber.Ctx) error {
	var params filter.Params
	if c.QueryParser(&params) != nil {
		return xerr.BadRequest(errors.New("invalid query parameters"))
	}

	errors := params.Validate()
	if len(errors) > 0 {
		return xerr.InvalidRequestData(errors)
	}

	ids := algo.AllProducts.ApplyFilter(&params)
	products := make([]product.Product, len(ids))
	for i, id := range ids {
		products[i] = product.Products[id]
	}

	return c.Status(fiber.StatusOK).JSON(products)
}
