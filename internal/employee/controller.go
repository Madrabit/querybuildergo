package employee

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"querybuilder/internal/web"
)

type Controller struct {
	svc    Svc
	server *web.Server
}

func NewController(server *web.Server, svc Svc) *Controller {
	return &Controller{server: server, svc: svc}
}

type Svc interface {
	FindByProducts(products []string) ([]byte, error)
}

func (c *Controller) RegisterRoutes() {
	c.server.R.Route("/employee", func(r chi.Router) {
		r.Post("/download", c.getFileByProducts)
	})
}

func (c *Controller) getFileByProducts(w http.ResponseWriter, r *http.Request) {
	var prodReq ProductsReq
	if err := json.NewDecoder(r.Body).Decode(&prodReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	products, err := c.svc.FindByProducts(prodReq.Products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileName := "emp.xls"
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(products)))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(products); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
