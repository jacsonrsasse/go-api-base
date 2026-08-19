package app

import (
	"github.com/go-chi/chi/v5"
	repository "github.com/jacsonrsasse/go-api-base/internal/adapters/postgresql/sqlc"
	"github.com/jacsonrsasse/go-api-base/internal/domain/something"
)

func addRoutes(h *chi.Mux, app *Application) error {

	somethingService := something.NewService(repository.New(app.Db))
	somethingHandler := something.NewHandler(somethingService)
	h.Get("/something", somethingHandler.ListSomethingMethod)
	return nil
}