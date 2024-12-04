package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/garrettladley/prods/internal/model/applicant"
	"github.com/garrettladley/prods/internal/xerr"
)

func (db *DB) Name(ctx context.Context, email applicant.NUEmail) (applicant.Name, error) {
	var name applicant.Name
	if err := db.db.GetContext(ctx, &name, "SELECT applicant_name FROM applicants WHERE email=$1;", email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", xerr.NotFound("applicant", "email", email)
		}
		return "", err
	}

	return name, nil
}
