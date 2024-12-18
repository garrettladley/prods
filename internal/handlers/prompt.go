package handlers

import (
	"fmt"

	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Service) Prompt(c *fiber.Ctx) error {
	rawToken := c.Params("token")
	token, err := uuid.Parse(rawToken)
	if err != nil {
		return xerr.BadRequest(fmt.Errorf("failed to parse token. got: %s", rawToken))
	}

	prompt, err := s.storage.Prompt(c.Context(), token)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(prompt)
}
