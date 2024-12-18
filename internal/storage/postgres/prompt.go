package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/xerr"
	go_json "github.com/goccy/go-json"
	"github.com/google/uuid"
)

func (db *DB) Prompt(ctx context.Context, token uuid.UUID) (algo.Prompt, error) {
	var r promptResult
	if err := db.db.GetContext(ctx, &r, "SELECT prompt FROM applicants WHERE token=$1;", token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return algo.Prompt{}, xerr.NotFound("prompt", "token", token)
		}
		return algo.Prompt{}, err
	}

	var productIDs []string
	if err := go_json.Unmarshal([]byte(r.Prompt.String), &productIDs); err != nil {
		return algo.Prompt{}, err
	}

	return algo.Prompt{
		ProductIDs: productIDs,
	}, nil
}

type promptResult struct {
	Prompt sql.NullString `db:"prompt"`
}
