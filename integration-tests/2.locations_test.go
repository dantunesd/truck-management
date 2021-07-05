package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLocations(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		method      string
		body        string
		wantStatus  int
		wantReponse string
	}{
		{
			"Should return 200 when getting the last location of an existing truck",
			"http://app:3000/trucks/1/locations/last",
			http.MethodGet,
			``,
			http.StatusOK,
			`{"id":3,"truck_id":1,"eld_id":"INSERTED 3","engine_state":"ON","current_speed":100,"latitude":1100,"longitude":1000,"engine_hours":1,"odometer":100,"created_at":"2021-07-04 23:45:56"}`,
		},
		{
			"Should return 200 when getting the last location of an existing truck without locations",
			"http://app:3000/trucks/3/locations/last",
			http.MethodGet,
			``,
			http.StatusOK,
			`{}`,
		},
		{
			"Should return 404 when getting a location for an inexistent truck",
			"http://app:3000/trucks/9999/locations/last",
			http.MethodGet,
			``,
			http.StatusNotFound,
			`{"title":"truck not found","status":404}`,
		},
		{
			"Should return 404 when getting a location for an invalid truck id",
			"http://app:3000/trucks/invalid/locations/last",
			http.MethodGet,
			``,
			http.StatusBadRequest,
			`{"title":"id must be numeric","status":400}`,
		},
		{
			"Should return 201 when creating a location with success",
			"http://app:3000/trucks/1/locations",
			http.MethodPost,
			`{"eld_id":"eld-id","engine_state":"ON","current_speed":100,"latitude":1100,"longitude":1000,"engine_hours":1,"odometer":100}`,
			http.StatusCreated,
			``,
		},
		{
			"Should return 404 when creating a location for an inexistent truck",
			"http://app:3000/trucks/9999/locations",
			http.MethodPost,
			`{"eld_id":"eld-id","engine_state":"ON","current_speed":100,"latitude":1100,"longitude":1000,"engine_hours":1,"odometer":100}`,
			http.StatusNotFound,
			`{"title":"truck not found","status":404}`,
		},
		{
			"Should return 400 when creating a location with an invalid truck id",
			"http://app:3000/trucks/invalid/locations",
			http.MethodPost,
			`{"eld_id":"eld-id","engine_state":"ON","current_speed":100,"latitude":1100,"longitude":1000,"engine_hours":1,"odometer":100}`,
			http.StatusBadRequest,
			`{"title":"id must be numeric","status":400}`,
		},
		{
			"Should return 400 when creating a location without required fields",
			"http://app:3000/trucks/1/locations",
			http.MethodPost,
			`{}`,
			http.StatusBadRequest,
			`{"title":"Key: 'CreateLocationInput.EldID' Error:Field validation for 'EldID' failed on the 'required' tag\nKey: 'CreateLocationInput.EngineState' Error:Field validation for 'EngineState' failed on the 'required' tag\nKey: 'CreateLocationInput.CurrentSpeed' Error:Field validation for 'CurrentSpeed' failed on the 'required' tag\nKey: 'CreateLocationInput.Latitude' Error:Field validation for 'Latitude' failed on the 'required' tag\nKey: 'CreateLocationInput.Longitude' Error:Field validation for 'Longitude' failed on the 'required' tag\nKey: 'CreateLocationInput.EngineHours' Error:Field validation for 'EngineHours' failed on the 'required' tag\nKey: 'CreateLocationInput.Odometer' Error:Field validation for 'Odometer' failed on the 'required' tag","status":400}`,
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
