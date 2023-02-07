package presentation_task

import (
	"matching-timestamps/domain"
	"net/http"
)

type PeriodicTaskHandler struct {
	pts domain.PeriodicTaskService
}

func (h PeriodicTaskHandler) PeriodicTaskList(w http.ResponseWriter, r *http.Request) {
	return
}

func NewPeriodicTaskHandler(pts domain.PeriodicTaskService) PeriodicTaskHandler {
	return PeriodicTaskHandler{pts: pts}
}
