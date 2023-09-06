package sqlite

import (
	"context"
	"errors"
	"fmt"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/migrations"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type Storage struct {
	DB *sqlx.DB
	mu sync.RWMutex
}

func NewStorage(path string) (*Storage, error) {
	const op = "adapter.database.sqlite.NewStorage"

	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{DB: db}, nil
}

func (s *Storage) RLock() {
	s.mu.RLock()
}

func (s *Storage) RUnlock() {
	s.mu.RUnlock()
}

func (s *Storage) Lock() {
	s.mu.Lock()
}

func (s *Storage) Unlock() {
	s.mu.Unlock()
}

func (s *Storage) RunMigrations() error {
	const op = "adapter.database.sqlite.RunMigrations"

	driver, err := sqlite3.WithInstance(s.DB.DB, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := bindata.Resource(migrations.AssetNames(), func(name string) ([]byte, error) {
		return migrations.Asset(name)
	})

	d, err := bindata.WithInstance(res)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	m, err := migrate.NewWithInstance("go-bindata", d, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	const op = "adapter.database.sqlite.Close"
	if err := s.DB.Close(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
