package product

import (
	"errors"
	"net/http"
	"querybuilder/internal/common"
)

type Handler struct {
	svc Svc
}

type Svc interface {
	GetAllProducts() (products Response, err error)
}

func NewHandler(svc Svc) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := h.svc.GetAllProducts()
	var nfErr *common.NotFoundError
	if err != nil {
		if errors.As(err, &nfErr) {
			common.ErrResponse(w, http.StatusOK, err.Error())
			return
		}
		common.ErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.OkResponse(w, products)
}
