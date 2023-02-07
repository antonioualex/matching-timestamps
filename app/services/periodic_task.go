package services

type PeriodicTaskService struct {
}

func (s PeriodicTaskService) MatchTimestamps(period, timeZone, startingPoint, endingPoint string) error {
	return nil
}

func NewPeriodicTaskService() *PeriodicTaskService {
	return &PeriodicTaskService{}
}
