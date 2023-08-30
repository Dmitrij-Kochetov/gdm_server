package main

import (
	"context"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/config"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/sqlite"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/logs"
	"log/slog"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	logger := logs.SetUpLogger(cfg.Env)

	logger = logger.With(slog.String("env", cfg.Env))
	logger.Info("Initializing server on:", slog.String("address", cfg.Address))
	logger.Debug("Debug mode enabled...")

	storage, err := sqlite.NewStorage(cfg.StoragePath)
	if err != nil {
		logger.Error("Storage initializing failed", logs.Err(err))
		os.Exit(1)
	}
	defer func(storage *sqlite.Storage, ctx context.Context) {
		err := storage.Close(ctx)
		if err != nil {
			logger.Error("Storage closing failed", logs.Err(err))
		}
	}(storage, context.Background())

	if err := storage.RunMigrations(); err != nil {
		logger.Error("Migrations failed", logs.Err(err))
		os.Exit(1)
	}

	//TODO: Init server: go-chi maybe

	//TODO: Run server
}
