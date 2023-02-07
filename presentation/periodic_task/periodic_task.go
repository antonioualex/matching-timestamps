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

type RequestedPeriodicTask struct {
	Period          string
	InvocationPoint domain.InvocationPoint
	Timezone        string
}

func (h PeriodicTaskHandler) MatchTimestamps(w http.ResponseWriter, r *http.Request) {

	period := r.URL.Query().Get("period")
	timeZone := r.URL.Query().Get("tz")
	startingPoint := r.URL.Query().Get("t1")
	endingPoint := r.URL.Query().Get("t2")

	if period == "" || timeZone == "" || startingPoint == "" || endingPoint == "" || len(startingPoint) != 16 || len(endingPoint) != 16 {
		w.WriteHeader(http.StatusBadRequest)
		badRequestResponse, _ := json.Marshal(BadRequestResponse{
			Status: "error",
			Desc:   "invalid parameters",
		})
		w.Write(badRequestResponse)
		return
	}

	requestedPeriodicTask := RequestedPeriodicTask{
		Period: period,
		InvocationPoint: domain.InvocationPoint{
			StartingTimestamp: startingPoint,
			EndingTimestamp:   endingPoint,
		},
		Timezone: timeZone,
	}

	timestamps, err := h.pts.GetPeriodicTimestamps(domain.PeriodicTask(requestedPeriodicTask))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		badRequestResponse, _ := json.Marshal(BadRequestResponse{
			Status: "error",
			Desc:   err.Error(),
		})
		w.Write(badRequestResponse)
		return
	}
	w.WriteHeader(http.StatusOK)
	timestampsResp, err := json.Marshal(timestamps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		badRequestResponse, _ := json.Marshal(BadRequestResponse{
			Status: "error",
			Desc:   "missing parameters",
		})
		w.Write(badRequestResponse)
		return
	}
	w.Write(timestampsResp)
}

func NewPeriodicTaskHandler(pts domain.PeriodicTaskService) PeriodicTaskHandler {
	return PeriodicTaskHandler{pts: pts}
}
