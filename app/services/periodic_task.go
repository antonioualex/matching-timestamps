package services

import (
	"errors"
	"log"
	"matching-timestamps/app/utils"
	"matching-timestamps/domain"
)

type PeriodicTaskService struct{}

func (s PeriodicTaskService) GetPeriodicTimestamps(pt domain.PeriodicTask) ([]string, error) {

	startingPointTimestamp, err := utils.StringRFC3339TimestampToTime(pt.InvocationPoint.StartingTimestamp)
	if err != nil {
		log.Println(err)
		return []string{""}, err
	}

	endingPointTimestamp, err := utils.StringRFC3339TimestampToTime(pt.InvocationPoint.EndingTimestamp)
	if err != nil {
		return []string{""}, err
	}

	if startingPointTimestamp.After(endingPointTimestamp) {
		return []string{""}, errors.New("starting point cannot be bigger than ending point")
	}

	locatedStartingPointTimestamp, err := utils.UTCtoLocation(startingPointTimestamp, pt.Timezone)
	if err != nil {
		return []string{""}, err
	}

	locatedEndingPointTimestamp, err := utils.UTCtoLocation(endingPointTimestamp, pt.Timezone)
	if err != nil {
		return []string{""}, err
	}

	timestamps, err := utils.GetTimestamps(pt.Period, locatedStartingPointTimestamp, locatedEndingPointTimestamp)
	if err != nil {
		return []string{""}, err
	}

	return timestamps, nil
}

func NewPeriodicTaskService() *PeriodicTaskService {
	return &PeriodicTaskService{}
}
