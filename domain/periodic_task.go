package domain

type PeriodicTask struct {
	Period          string
	InvocationPoint InvocationPoint
	Timezone        string
}

type InvocationPoint struct {
	StartingTimestamp string
	EndingTimestamp   string
}

type PeriodicTaskService interface {
}
