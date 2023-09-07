package migrate

import (
	"errors"
	"fmt"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/migrations"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/jmoiron/sqlx"
)

type Migrator struct {
	migrate *migrate.Migrate
}

func NewMigrator(db *sqlx.DB) (*Migrator, error) {
	const op = "adapter.database.sqlite.migrate.NewMigrator"

	driver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	res := bindata.Resource(migrations.AssetNames(), func(name string) ([]byte, error) {
		return migrations.Asset(name)
	})

	d, err := bindata.WithInstance(res)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	m, err := migrate.NewWithInstance("go-bindata", d, "sqlite3", driver)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Migrator{migrate: m}, nil
}

func (m *Migrator) RunUp() error {
	const op = "adapter.database.sqlite.migrate.RunUp"

	err := m.migrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (m *Migrator) RunDown() error {
	const op = "adapter.database.sqlite.migrate.RunDown"

	err := m.migrate.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
