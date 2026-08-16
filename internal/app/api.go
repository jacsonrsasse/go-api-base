package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config Config
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return r
}

func (app *Application) Run(h http.Handler) error {
	svr := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      h,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server started on %s", app.Config.Addr)

	return svr.ListenAndServe()
}

type Config struct {
	Addr string
	DB DbConfig
}

type DbConfig struct {
	Dsn string
}
