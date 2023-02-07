package periodic_task

import (
	"errors"
	"matching-timestamps/app/fakes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPeriodicTaskHandler_MatchTimestamps(t *testing.T) {
	srv := &fakes.FakePeriodicTaskService{}

	results := []string{
		"20180214T204603Z",
		"20190214T204603Z",
		"20200214T204603Z",
	}
	srv.GetPeriodicTimestampsReturns(results, nil)

	myHandler := PeriodicTaskHandler{srv}
	ts := httptest.NewServer(http.HandlerFunc(myHandler.MatchTimestamps))

	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL+"?period=1y&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z", nil)
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is not 201, but is %v", res.StatusCode)
	}

}

func TestPeriodicTaskHandler_MatchTimestamps_FAIL(t *testing.T) {
	srv := &fakes.FakePeriodicTaskService{}

	results := []string{""}
	srv.GetPeriodicTimestampsReturns(results, errors.New("not reachable"))

	myHandler := PeriodicTaskHandler{srv}
	ts := httptest.NewServer(http.HandlerFunc(myHandler.MatchTimestamps))

	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is not 400, but is %v", res.StatusCode)
	}

}
