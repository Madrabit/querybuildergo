package employee

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	filepath "path/filepath"
	"time"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) getFileByProducts(w http.ResponseWriter, r *http.Request) {
	var prodReq ProductsReq
	if err := json.NewDecoder(r.Body).Decode(&prodReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	products, err := h.store.FindByProducts(r.Context(), prodReq.Products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = CreateExl(products)
	if err != nil {
		return
	}
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	downloadsDir := filepath.Join(rootDir, "downloads")
	filePath := filepath.Join(downloadsDir, "emp.xlsx")
	f, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	fileName := filepath.Base(filePath)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	http.ServeContent(w, r, fileName, time.Now(), f)
}
