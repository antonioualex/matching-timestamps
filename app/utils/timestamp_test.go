package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimestamps(t *testing.T) {
	period := "1h"
	locatedStartingPointTimestamp := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	locatedEndingPointTimestamp := time.Date(2009, 1, 1, 7, 0, 0, 0, time.UTC)
	timestamps, err := GetTimestamps(period, locatedStartingPointTimestamp, locatedEndingPointTimestamp)
	if err != nil {
		t.Error(err)
	}

	if len(timestamps) != 7 {
		t.Error(fmt.Sprintf("output number of timestamps should be 7, but is %d", len(timestamps)))
	}
}

func TestGetTimestamps_FAIL(t *testing.T) {
	period := "1g"
	locatedStartingPointTimestamp := time.Date(2009, 1, 1, 7, 0, 0, 0, time.UTC)
	locatedEndingPointTimestamp := time.Date(2009, 1, 1, 7, 0, 0, 0, time.UTC)

	_, err := GetTimestamps(period, locatedStartingPointTimestamp, locatedEndingPointTimestamp)
	if err == nil {
		t.Error(err)
	}

}

func TestStringRFC3339TimestampToTime(t *testing.T) {

	timestampStr := "20060102T150405Z"
	formattedTimestamp, err := StringRFC3339TimestampToTime(timestampStr)
	if err != nil {
		t.Error(err)
	}

	if formattedTimestamp.Year() != 2006 {
		t.Error("formatted timestamp year does not match")
	}
	if formattedTimestamp.Month() != 1 {
		t.Error("formatted timestamp month does not match")
	}

	if formattedTimestamp.Day() != 2 {
		t.Error("formatted timestamp day does not match")
	}

}

func TestStringRFC3339TimestampToTime_FAIL(t *testing.T) {
	timestampStr := "1060102T150405Z"
	_, err := StringRFC3339TimestampToTime(timestampStr)
	if err == nil {
		t.Error(err)
	}

}

func TestUTCtoLocation(t *testing.T) {
	locationStr := "Europe/Athens"
	timestamp := time.Date(2009, 1, 1, 7, 0, 0, 0, time.UTC)
	convertedTimestamp, err := UTCtoLocation(timestamp, locationStr)
	if err != nil {
		t.Error(err)
	}

	if convertedTimestamp.Hour() != timestamp.Hour() {
		t.Error("converted timestamp hour shouldn't match with the original one")
	}
}

func TestUTCtoLocation_FAIL(t *testing.T) {
	locationStr := "Europ/Athens"
	timestamp := time.Date(2009, 1, 1, 7, 0, 0, 0, time.UTC)
	_, err := UTCtoLocation(timestamp, locationStr)
	if err == nil {
		t.Error(err)
	}
}
