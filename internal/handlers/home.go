package handlers

import (
	"github.com/garrettladley/prods/internal/views/home"
	"github.com/garrettladley/prods/internal/xtempl"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Home(c *fiber.Ctx) error {
	return xtempl.Render(c, home.Index())
}
