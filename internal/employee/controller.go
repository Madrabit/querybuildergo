package employee

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"log"
	"net/http"
	"querybuilder/internal/common"
	"querybuilder/internal/web"
)

type Controller struct {
	svc    Svc
	server *web.Server
	logger *common.Logger
}

func NewController(server *web.Server, svc Svc, logger *common.Logger) *Controller {
	return &Controller{server: server, svc: svc, logger: logger}
}

type Svc interface {
	FindByProducts(products []string) ([]byte, error)
}

func (c *Controller) RegisterRoutes() {
	c.server.R.Route("/employee", func(r chi.Router) {
		r.Post("/download", c.getFileByProducts)
	})
}

// getFileByProducts godoc
// @Summary      Получить Excel-файл по продуктам
// @Description  Возвращает Excel-файл с данными по списку продуктов
// @Tags         employees
// @Accept       json
// @Produce      application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param        request body ProductsReq true "Список продуктов"
// @Success      200 {file} file "Файл с данными"
// @Failure      400 {string} string "bad request"
// @Router       /employee/download [post]
func (c *Controller) getFileByProducts(w http.ResponseWriter, r *http.Request) {
	var request ProductsReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.logger.Error("failed to get employees", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.logger.Debug("get employees: received request", zap.Any("request", request))
	products, err := c.svc.FindByProducts(request.Products)
	if err != nil {
		c.logger.Error("failed to get employees", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileName := "emp.xls"
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(products)))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(products); err != nil {
		c.logger.Error("failed to write response", zap.Error(err))
		log.Printf("failed to write response: %v", err)
	}
	c.logger.Info("successfully retrieve employees list")
}
