package manager

import (
	"encoding/json"
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

func NewController(server *web.Server, svc Svc) *Controller {
	return &Controller{server: server, svc: svc}
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
	var report DailyReportDTOReq
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dailyReport, err := c.svc.GetDailyReport(report.Manager, report.StartDate, report.EndDate)
	var reqErr *common.RequestValidationError
	var notFoundErr *common.NotFoundError
	if err != nil {
		switch {
		case errors.As(err, &notFoundErr):
			common.OkResponseMsg(w, dailyReport, "controller manager: get daily report: report not found")
			return
		case errors.As(err, &reqErr):
			common.ErrResponse(w, http.StatusBadRequest, err.Error())
			return
		default:
			common.ErrResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	common.OkResponse(w, dailyReport)
}
