package handlers

import (
	"fmt"

	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Prompt godoc
//
//	@Summary		Retrieve registration challenge prompt
//	@Description	Fetches the challenge prompt associated with a specific registration token.
//	@Tags			applicants
//	@Produce		json
//	@Param			token	path		string			true	"Registration Token"	format(uuid)
//	@Success		200		{object}	algo.Prompt		"Successfully retrieved challenge prompt"
//	@Failure		400		{object}	xerr.APIError	"Invalid token format"
//	@Failure		404		{object}	xerr.APIError	"Token not found"
//	@Failure		429		{object}	xerr.APIError	"Too many requests"
//	@Failure		500		{object}	xerr.APIError	"Internal server error"
//	@Router			/api/v1/{token}/prompt [get]
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
