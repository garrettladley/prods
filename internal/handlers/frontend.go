package handlers

import (
	"github.com/garrettladley/prods/internal/views/challenges/frontend"
	"github.com/garrettladley/prods/internal/xtempl"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Frontend(c *fiber.Ctx) error {
	return xtempl.Render(c, frontend.Index())
}
