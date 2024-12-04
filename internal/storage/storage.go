package storage

import (
	"context"

	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/google/uuid"
)

type Storage interface {
	Token(ctx context.Context, email applicant.NUEmail) (uuid.UUID, error)
	Name(ctx context.Context, email applicant.NUEmail) (applicant.Name, error)
}
