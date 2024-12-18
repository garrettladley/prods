package handlers

import (
	"fmt"
	"net/url"

	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Token(c *fiber.Ctx) error {
	rawEmail := c.Query("email")
	unescapedEmail, err := url.QueryUnescape(rawEmail)
	if err != nil {
		return xerr.BadRequest(fmt.Errorf("failed to unescape email. got: %s", rawEmail))
	}
	email, err := applicant.ParseNUEmail(unescapedEmail)
	if err != nil {
		return xerr.BadRequest(fmt.Errorf("failed to parse email. got: %s", email))
	}
	token, err := s.storage.Token(c.Context(), email)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).SendString(token.String())
}
