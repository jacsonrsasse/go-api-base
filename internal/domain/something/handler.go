package something

import (
	"net/http"

	"github.com/jacsonrsasse/go-api-base/internal/helpers/response"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListSomethingMethod(w http.ResponseWriter, r *http.Request) {
	something, err := h.service.ListSomething(r.Context())
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err)
		return
	}
	response.Send(w, http.StatusOK, something)
}
