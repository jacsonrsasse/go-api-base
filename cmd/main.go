package main

import (
	"log"
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

	if err := app.Run(app.Mount()); err != nil {
		log.Printf("server has failed to start, err %s", err)
		os.Exit(1)
	}
}
