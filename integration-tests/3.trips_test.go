package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTrips(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		method      string
		body        string
		wantStatus  int
		wantReponse string
	}{
		{
			"Should return 200 when getting the trip summary of an existing truck",
			"http://app:3000/trucks/4/trips/summary",
			http.MethodGet,
			``,
			http.StatusOK,
			`{"id":1,"truck_id":4,"origin":"90 80","destination":"90 80","state":"ONGOING","odometer":100,"engine_hours":5,"average_speed":100,"updated_at":"2021-07-05 00:41:37"}`,
		},
		{
			"Should return 200 when getting the trip summary of an existing truck without trip",
			"http://app:3000/trucks/3/trips/summary",
			http.MethodGet,
			``,
			http.StatusOK,
			`{}`,
		},
		{
			"Should return 404 when getting a trip summary for an inexistent truck",
			"http://app:3000/trucks/9999/trips/summary",
			http.MethodGet,
			``,
			http.StatusNotFound,
			`{"title":"truck not found","status":404}`,
		},
		{
			"Should return 404 when getting a trip summary with an invalid truck id",
			"http://app:3000/trucks/invalid/trips/summary",
			http.MethodGet,
			``,
			http.StatusBadRequest,
			`{"title":"id must be numeric","status":400}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var content io.Reader
			if len(tt.body) > 0 {
				content = bytes.NewBuffer([]byte(tt.body))
			}

			req, _ := http.NewRequest(tt.method, tt.url, content)

			h := http.Client{}
			resp, err := h.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if status := resp.StatusCode; status != tt.wantStatus {
				t.Errorf("status code: got %v want %v",
					status, tt.wantStatus)
			}

			r, _ := ioutil.ReadAll(resp.Body)
			if tt.wantReponse != `` && string(r) != tt.wantReponse {
				t.Errorf("body: got %v want %v",
					string(r), tt.wantReponse)
			}
		})
	}
}
