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

func (db *DB) Solution(ctx context.Context, token uuid.UUID) (s algo.Solution, err error) {
	var r solutionResult
	if err := db.db.GetContext(ctx, &r, "SELECT solution FROM applicants WHERE token=$1;", token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s, xerr.NotFound("solution", "token", token)
		}
	}

	var orderedProductIDs [][]string
	if err := go_json.Unmarshal([]byte(r.Solution.String), &orderedProductIDs); err != nil {
		return s, err
	}

	s.OrderedProductIDs = orderedProductIDs

	return
}

type solutionResult struct {
	Solution sql.NullString `db:"solution"`
}
