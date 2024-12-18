package postgres

import "github.com/lib/pq"

const uniqueViolationCode = "23505"

func (db *DB) isUniqueViolation(err error) bool {
	pgErr, isPGError := err.(*pq.Error)
	return isPGError && pgErr.Code == uniqueViolationCode
}
