package postgres

import (
	"fmt"
	"time"

	"github.com/garrettladley/prods/internal/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DSN             string
	MaxOpenConns    uint
	MaxIdleConns    uint
	ConnMaxLifetime time.Duration
}

func From(settings settings.DB) Config {
	return Config{
		DSN:             settings.DSN,
		MaxOpenConns:    settings.MaxOpenConns,
		MaxIdleConns:    settings.MaxIdleConns,
		ConnMaxLifetime: settings.ConnMaxLifetime,
	}
}

type DB struct {
	db *sqlx.DB
}

func New(cfg Config) (DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return DB{}, fmt.Errorf("error connecting to database: %v", err)
	}

	db.SetMaxOpenConns(int(cfg.MaxOpenConns))
	db.SetMaxIdleConns(int(cfg.MaxIdleConns))
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return DB{db}, nil
}
