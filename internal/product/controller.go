package product

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"querybuilder/internal/common"
	"querybuilder/internal/web"
)

type Controller struct {
	server *web.Server
	svc    Svc
}

type Svc interface {
	GetAllProducts() (products Response, err error)
}

func NewController(server *web.Server, svc Svc) *Controller {
	return &Controller{
		server: server,
		svc:    svc}
}

func (c *Controller) RegisterRoutes() {
	c.server.R.Route("/products", func(r chi.Router) {
		r.Get("/", c.GetProducts)
	})
}

func (c *Controller) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := c.svc.GetAllProducts()
	var nfErr *common.NotFoundError
	if err != nil {
		if errors.As(err, &nfErr) {
			common.OkResponseMsg(w, products, "controller product: get products: products not found")
			return
		}
		common.ErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.OkResponse(w, products)
}
