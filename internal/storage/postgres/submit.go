package postgres

import (
	"context"

	"github.com/google/uuid"
)

func (db *DB) Submit(ctx context.Context, token uuid.UUID, score int) error {
	_, err := db.db.ExecContext(
		ctx,
		"INSERT INTO submissions (submission_id, token, score) VALUES ($1, $2, $3);",
		uuid.New(),
		token,
		score,
	)

	return err
}
