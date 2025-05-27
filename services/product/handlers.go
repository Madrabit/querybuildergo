package product

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetProducts(w http.ResponseWriter, req *http.Request) {
	products, err := h.store.GetAllProducts(req.Context())
	pr := make([]string, 0)
	for _, product := range products {
		p := ToProductDto(product)
		pr = append(pr, p)
	}
	resp := ProductName{ProductName: pr}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
