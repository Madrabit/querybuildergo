package common

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"error"`
	Data    T      `json:"data"`
}

func ErrResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(Response[any]{
		Success: false,
		Message: msg,
		Data:    nil,
	})
}

func OkResponse[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response[any]{
		Success: true,
		Message: "",
		Data:    data,
	})
}
