package storage

import (
	"context"
	"time"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/google/uuid"
)

type Storage interface {
	Register(ctx context.Context, r Register) error
	Token(ctx context.Context, email applicant.NUEmail) (uuid.UUID, error)
	Prompt(ctx context.Context, token uuid.UUID) (algo.Prompt, error)
	Solution(ctx context.Context, token uuid.UUID) (algo.Solution, error)
	Submit(ctx context.Context, token uuid.UUID, score int) error
}

type Register struct {
	Email     applicant.NUEmail
	Name      applicant.Name
	CreatedAt time.Time
	Token     uuid.UUID
	Prompt    algo.Prompt
	Solution  algo.Solution
}
