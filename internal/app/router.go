package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jacsonrsasse/go-api-base/internal/domain/something"
)

func addRoutes(h *chi.Mux) error {

	somethingService := something.NewService()
	somethingHandler := something.NewHandler(somethingService)
	h.Get("/something", somethingHandler.ListSomethingMethod)
	return nil
}