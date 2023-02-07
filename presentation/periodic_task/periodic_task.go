package periodic_task

import (
	"encoding/json"
	"matching-timestamps/domain"
	"net/http"
)

type PeriodicTaskHandler struct {
	pts domain.PeriodicTaskService
}

type BadRequestResponse struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

func (h PeriodicTaskHandler) MatchTimestamps(w http.ResponseWriter, r *http.Request) {

	period := r.URL.Query().Get("period")
	timeZone := r.URL.Query().Get("tz")
	startingPoint := r.URL.Query().Get("pt1")
	endingPoint := r.URL.Query().Get("pt2")

	if period == "" || timeZone == "" || startingPoint == "" || endingPoint == "" {
		w.WriteHeader(http.StatusBadRequest)
		badRequestResponse, _ := json.Marshal(BadRequestResponse{
			Status: "error",
			Desc:   "missing parameters",
		})
		w.Write(badRequestResponse)
		return
	}

	err := h.pts.MatchTimestamps(period, timeZone, startingPoint, endingPoint)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		badRequestResponse, _ := json.Marshal(BadRequestResponse{
			Status: "error",
			Desc:   "failed to match timestamps",
		})
		w.Write(badRequestResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func NewPeriodicTaskHandler(pts domain.PeriodicTaskService) PeriodicTaskHandler {
	return PeriodicTaskHandler{pts: pts}
}
