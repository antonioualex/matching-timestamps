package presentation_task

import "matching-timestamps/domain"

func CreateRoutes(pts domain.PeriodicTaskService) map[string]domain.RouteDef {
	ph := NewPeriodicTaskHandler(pts)
	return map[string]domain.RouteDef{
		"/ptlist": {
			Methods:     []string{"GET"},
			HandlerFunc: ph.PeriodicTaskList,
		},
	}

}
