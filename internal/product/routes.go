package product

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(handler *Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", handler.GetProducts)
	return r
}
