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
			"Should return 201 when creating a location with success",
			"http://app:3000/trucks/1/locations",
			http.MethodPost,
			`{"eld_id":"eld-id","engine_state":"ON","current_speed":100,"latitude":1100,"longitude":1000,"engine_hours":1,"odometer":100}`,
			http.StatusCreated,
			``,
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
