package handlers

import (
	"github.com/garrettladley/prods/internal/model/product"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
)

// Product godoc
//
//	@Summary		Retrieve a product by ID
//	@Description	Fetches a specific product using its unique identifier
//	@Description	Note: Please store the your fetched products locally; whether it be in memory, on disk, or in a database.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Product ID"
//	@Success		200	{object}	product.Product	"Product details"
//
// @Failure		404	{object}	xerr.APIError	"Product not found"
// @Failure		429	{object}	xerr.APIError	"Too many requests"
//
// @Failure		500	{object}	xerr.APIError	"Internal server error"
// @Router			/api/v1/products/{id} [get]
func (s *Service) Product(c *fiber.Ctx) error {
	id := c.Params("id")
	product, ok := product.Products[id]
	if !ok {
		return xerr.NotFound("product", "id", id)
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
