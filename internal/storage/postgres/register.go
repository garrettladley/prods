package postgres

import (
	"context"

	"github.com/garrettladley/prods/internal/storage"
	"github.com/garrettladley/prods/internal/xerr"

	go_json "github.com/goccy/go-json"
)

func (db *DB) Register(ctx context.Context, r storage.Register) error {
	marshalledPrompt, err := go_json.Marshal(r.Prompt.ProductIDs)
	if err != nil {
		return err
	}

	marshalledRanking, err := go_json.Marshal(r.Solution.OrderedProductIDs)
	if err != nil {
		return err
	}

	if _, err := db.db.ExecContext(
		ctx,
		"INSERT INTO applicants (email, applicant_name, created_at, token, prompt, solution) VALUES ($1, $2, $3, $4, $5, $6);",
		r.Email,
		r.Name,
		r.CreatedAt,
		r.Token,
		marshalledPrompt,
		marshalledRanking,
	); err != nil {
		if db.isUniqueViolation(err) {
			return xerr.Conflict("user", "email", r.Email)
		} else {
			return err
		}
	}

	return nil
}
