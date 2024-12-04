package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/google/uuid"
)

type tokenResult struct {
	Token sql.NullString `db:"token"`
}

func (db *DB) Token(ctx context.Context, email applicant.NUEmail) (uuid.UUID, error) {
	var dbResult tokenResult
	if err := db.db.GetContext(ctx, &dbResult, "SELECT token FROM applicants WHERE email=$1;", email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.UUID{}, xerr.NotFound("applicant", "email", email)
		}
		return uuid.UUID{}, err
	}

	token, err := uuid.Parse(dbResult.Token.String)
	if err != nil {
		return uuid.UUID{}, err
	}

	return token, nil
}
