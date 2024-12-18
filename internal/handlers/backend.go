package handlers

import (
	"github.com/garrettladley/prods/internal/views/backend"
	"github.com/garrettladley/prods/internal/xtempl"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Backend(c *fiber.Ctx) error {
	return xtempl.Render(c, backend.Index())
}
