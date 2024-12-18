package handlers

import (
	"log/slog"
	"time"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/garrettladley/prods/internal/storage"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type registerRequest struct {
	RawName  string `json:"name"`
	RawEmail string `json:"email"`
}

type registerResponse struct {
	Token  uuid.UUID   `json:"token"`
	Prompt algo.Prompt `json:"prompt"`
}

// Register godoc
//
//	@Summary		Register a new applicant
//	@Description	Creates a new applicant registration with a unique token and challenge prompt.
//	@Description	Only accepts your @northeastern.edu email.
//	@Description	Note: Please store the challenge prompt locally; whether it be in memory, on disk, or in a database.
//	@Tags			applicants
//	@Accept			json
//	@Produce		json
//	@Param			request	body		registerRequest		true	"Registration Request"
//	@Success		201		{object}	registerResponse	"Successfully registered applicant"
//	@Failure		400		{object}	xerr.APIError		"Invalid JSON"
//	@Failure		409		{object}	xerr.APIError		"Email conflict"
//	@Failure		422		{object}	xerr.APIError		"Unprocessable entity"
//	@Failure		429		{object}	xerr.APIError		"Too many requests"
//	@Failure		500		{object}	xerr.APIError		"Internal server error"
//	@Router			/api/v1/register [post]
func (s *Service) Register(c *fiber.Ctx) error {
	var r registerRequest
	if err := c.BodyParser(&r); err != nil {
		slog.Error("invalid JSON request data", "error", err)
		return xerr.InvalidJSON()
	}

	validated, errors := r.validate()
	if len(errors) > 0 {
		return xerr.InvalidRequestData(errors)
	}

	var (
		token  = uuid.New()
		now    = time.Now()
		prompt = s.algo.Generate(uint64(now.UnixNano()))
	)

	if err := s.storage.Register(
		c.Context(),
		storage.Register{
			Email:     validated.Email,
			Name:      validated.Name,
			CreatedAt: now,
			Token:     token,
			Prompt:    *prompt,
			Solution:  s.algo.Solution(c.Context(), *prompt),
		},
	); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(registerResponse{Token: token, Prompt: *prompt})
}

type validatedRegisterRequest struct {
	Email applicant.NUEmail
	Name  applicant.Name
}

func (r *registerRequest) validate() (validatedRegisterRequest, map[string]string) {
	errors := make(map[string]string)
	email, err := applicant.ParseNUEmail(r.RawEmail)
	if err != nil {
		errors["email"] = err.Error()
	}

	name, err := applicant.ParseName(r.RawName)
	if err != nil {
		errors["name"] = err.Error()
	}

	return validatedRegisterRequest{
		Email: email,
		Name:  name,
	}, errors
}
