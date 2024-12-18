package handlers

import (
	"fmt"
	"net/url"

	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
)

// Token godoc
//
//	@Summary		Retrieve registration token
//	@Description	Fetches the token associated with a specific email.
//
// @Description Note: The email must be query-escaped. Replace `@` with `%40` and other special characters as per URL encoding standards.
//
//	@Tags			applicants
//	@Produce		plain
//	@Param			email	query		string			true	"Email"	format(email)
//	@Success		200		{string}	string			"Successfully retrieved token"
//	@Failure		400		{object}	xerr.APIError	"Invalid email format"
//	@Failure		404		{object}	xerr.APIError	"Email not found"
//	@Failure		429		{object}	xerr.APIError	"Too many requests"
//	@Failure		500		{object}	xerr.APIError	"Internal server error"
//	@Router			/api/v1/token/{email} [get]
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
