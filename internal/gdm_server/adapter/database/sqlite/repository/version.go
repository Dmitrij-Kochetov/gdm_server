package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VersionRepository struct {
	db *sqlx.DB
}

func NewVersionRepo(path string) (*VersionRepository, error) {
	const op = "adapter.database.sqlite.repository.version"

	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &VersionRepository{db: db}, nil
}

func (v *VersionRepository) Close(ctx context.Context) error {
	const op = "adapter.database.sqlite.repository.version"
	if err := v.db.Close(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
