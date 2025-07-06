package product

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"querybuilder/internal/common"
	"querybuilder/internal/web"
)

type Controller struct {
	server *web.Server
	svc    Svc
	logger *common.Logger
}

type Svc interface {
	GetAllProducts() (products Response, err error)
}

func NewController(server *web.Server, svc Svc, logger *common.Logger) *Controller {
	return &Controller{
		server: server,
		svc:    svc,
		logger: logger,
	}
}

func (c *Controller) RegisterRoutes() {
	c.server.R.Route("/products", func(r chi.Router) {
		r.Get("/", c.GetProducts)
	})
}

// GetProducts godoc
// @Summary      Получить список продуктов
// @Description  Возвращает полный список доступных продуктов
// @Tags         products
// @Produce      json
// @Success      200 {array} Response "Список продуктов"
// @Failure      500 {string} string "Внутренняя ошибка сервера"
// @Router       /products/ [get]
func (c *Controller) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := c.svc.GetAllProducts()
	var nfErr *common.NotFoundError
	if err != nil {
		c.logger.Error("failed to  get products", zap.Error(err))
		if errors.As(err, &nfErr) {
			c.logger.Error("products not found", zap.Error(err))
			common.OkResponseMsg(w, products, "controller product: get products: products not found")
			return
		}
		common.ErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.logger.Info("products retrieved successfully")
	common.OkResponse(w, products)
}
