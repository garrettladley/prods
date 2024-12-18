package handlers

import (
	"errors"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/filter"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
)

// TODO: document me
func (s *Service) Products(c *fiber.Ctx) error {
	var params filter.Params
	if c.QueryParser(&params) != nil {
		return xerr.BadRequest(errors.New("invalid query parameters"))
	}

	errors := params.Validate()
	if len(errors) > 0 {
		return xerr.InvalidRequestData(errors)
	}

	if params.Limit == 0 {
		params.Limit = 3
	}

	return c.Status(fiber.StatusOK).JSON(algo.AllProducts.ApplyFilter(&params))
}
