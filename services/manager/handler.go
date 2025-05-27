package manager

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

func (h *Handler) getDailyReport(w http.ResponseWriter, r *http.Request) {
	var report DailyReportDTOReq
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	dailyReport, err := h.store.GetDailyReport(r.Context(), report.Manager, report.StartDate, report.EndDate)
	var rep []CallReportMapped
	for _, r := range dailyReport {
		callReport := r.ToCallReport()
		rep = append(rep, callReport)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	resp := ManagerReportResp{ManagerReport: rep}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
