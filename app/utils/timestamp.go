package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func stringRFC3339ToCustomRFC3339(timestamp string) (string, error) {
	pattern := `(.{4})(.{1})(.{2})(.{1})(.{5})(.{1})(.{2})(.{1})(.{3})`
	result, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return "", err
	}
	timestampResult := result.ReplaceAllString(timestamp, "$1$3$5$7$9")

	return timestampResult, nil
}

func StringRFC3339TimestampToTime(timestampStr string) (time.Time, error) {
	pattern := `(.{4})(.{2})(.{2})(.{1})(.{2})(.{2})(.{2})(.{1})`
	result, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return time.Time{}, err
	}
	timestampStr = result.ReplaceAllString(timestampStr, "$1-$2-$3$4$5:$6:$7$8")

	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		log.Println(err)
		return time.Time{}, err
	}
	return timestamp, nil
}

func UTCtoLocation(timestamp time.Time, location string) (time.Time, error) {
	locationByName, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}

	return timestamp.In(locationByName).UTC(), nil
}

func GetTimestamps(period string, locatedStartingPointTimestamp, locatedEndingPointTimestamp time.Time) ([]string, error) {

	period, duration, err := splitPeriod(period)
	if err != nil {
		return []string{""}, err
	}

	dur := time.Duration(duration)
	var results []string
	switch period {
	case "h":
		for t := locatedStartingPointTimestamp; t.Before(locatedEndingPointTimestamp) || t.Equal(locatedEndingPointTimestamp); t = t.Add(time.Hour * dur) {
			if t.Hour() != locatedEndingPointTimestamp.Hour() {
				result, err := stringRFC3339ToCustomRFC3339(t.Format(time.RFC3339))
				if err != nil {
					log.Println(err)
				}
				results = append(results, result)
			}
		}
	case "d":
		for t := locatedStartingPointTimestamp; t.Before(locatedEndingPointTimestamp); t = t.AddDate(0, 0, duration) {
			if t.Day() != locatedEndingPointTimestamp.Day() {
				result, err := stringRFC3339ToCustomRFC3339(t.Format(time.RFC3339))
				if err != nil {
					log.Println(err)
				}
				results = append(results, result)
			}
		}
	case "mo":
		for t := locatedStartingPointTimestamp; t.Before(locatedEndingPointTimestamp); t = t.AddDate(0, duration, 0) {
			if t.Month() != locatedEndingPointTimestamp.Month() {
				result, err := stringRFC3339ToCustomRFC3339(t.Format(time.RFC3339))
				if err != nil {
					log.Println(err)
				}
				results = append(results, result)
			}
		}
	case "y":
		for t := locatedStartingPointTimestamp; t.Before(locatedEndingPointTimestamp); t = t.AddDate(duration, 0, 0) {
			if t.Year() != locatedEndingPointTimestamp.Year() {
				result, err := stringRFC3339ToCustomRFC3339(t.Format(time.RFC3339))
				if err != nil {
					log.Println(err)
				}
				results = append(results, result)
			}
		}
	default:
		return []string{""}, errors.New("invalid period parameter")
	}
	return results, nil
}

func splitPeriod(period string) (string, int, error) {

	re := regexp.MustCompile(`(\d+)(\w{1,2})`)
	matches := re.FindStringSubmatch(period)

	amount, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", 0, errors.New("invalid period parameter")
	}

	unit := matches[2]

	return unit, amount, nil
}
