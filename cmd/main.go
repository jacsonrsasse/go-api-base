package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jacsonrsasse/go-api-base/internal/app"
	"github.com/jacsonrsasse/go-api-base/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := app.Config{
		Addr: ":8080",
		DB:   app.DbConfig{
			Dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=something_db sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.DB.Dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to the database")

	app := app.Application{
		Config: cfg,
	}

	if err := app.Run(app.Mount()); err != nil {
		slog.Error("server has failed to start", "err", err)
		os.Exit(1)
	}
}
