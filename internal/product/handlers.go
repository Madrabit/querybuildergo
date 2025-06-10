package product

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	store *Repository
}

func NewHandler(store *Repository) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetProducts(w http.ResponseWriter, req *http.Request) {
	products, err := h.store.GetAllProducts(req.Context())
	response := ToResponse(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
