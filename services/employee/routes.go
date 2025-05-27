package employee

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(handler *Handler) http.Handler {
	r := chi.NewRouter()
	r.Post("/download", handler.getFileByProducts)
	return r
}
