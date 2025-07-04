package manager

import (
	"encoding/json"
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

func NewController(server *web.Server, svc Svc, logger *common.Logger) *Controller {
	return &Controller{server: server, svc: svc, logger: logger}
}

type Svc interface {
	GetDailyReport(manager, startDate, endDate string) (Response, error)
}

func (c *Controller) RegisterRoutes() {
	c.server.R.Route("/managers", func(r chi.Router) {
		r.Post("/", c.GetDailyReport)
	})
}

func (c *Controller) GetDailyReport(w http.ResponseWriter, r *http.Request) {
	var request DailyReportDTOReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.logger.Debug("get daily request: received request", zap.Any("request", request))
	dailyReport, err := c.svc.GetDailyReport(request.Manager, request.StartDate, request.EndDate)
	var reqErr *common.RequestValidationError
	var notFoundErr *common.NotFoundError
	if err != nil {
		c.logger.Error("failed to get daily request", zap.Error(err))
		switch {
		case errors.As(err, &notFoundErr):
			c.logger.Error("daily request not found", zap.Error(err))
			common.OkResponseMsg(w, dailyReport, "controller manager: get daily request: request not found")
			return
		case errors.As(err, &reqErr):
			common.ErrResponse(w, http.StatusBadRequest, err.Error())
			return
		default:
			common.ErrResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.logger.Info("successfully retrieve daily request")
	common.OkResponse(w, dailyReport)
}
