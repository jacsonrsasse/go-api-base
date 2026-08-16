package main

import (
	"log/slog"
	"os"

	"github.com/jacsonrsasse/go-api-base/internal/app"
)

func main() {
	cfg := app.Config{
		Addr: ":8080",
		DB:   app.DbConfig{},
	}

	app := app.Application{
		Config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := app.Run(app.Mount()); err != nil {
		slog.Error("server has failed to start", "err", err)
		os.Exit(1)
	}
}
