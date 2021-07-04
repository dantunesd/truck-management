package tests

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestTrucks(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		body       string
		wantStatus int
	}{
		{
			"Should return 201 when creating a truck with success",
			"http://app:3000/trucks",
			http.MethodPost,
			`{"license_plate":"ABC1234","eld_id":"eld-id","carrier_id":"My Carrier","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900}`,
			http.StatusCreated,
		},
		{
			"Should return 409 when creating a truck with with the same license plate",
			"http://app:3000/trucks",
			http.MethodPost,
			`{"license_plate":"ABC1234","eld_id":"new-eld-id","carrier_id":"My Carrier","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900}`,
			http.StatusConflict,
		},
		{
			"Should return 409 when creating a truck with with the same eld id",
			"http://app:3000/trucks",
			http.MethodPost,
			`{"license_plate":"ABC4567","eld_id":"eld-id","carrier_id":"My Carrier","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900}`,
			http.StatusConflict,
		},
		{
			"Should return 400 when creating a truck without required fields",
			"http://app:3000/trucks",
			http.MethodPost,
			`{}`,
			http.StatusBadRequest,
		},
		{
			"Should return 400 when creating a truck with invalid payload",
			"http://app:3000/trucks",
			http.MethodPost,
			`invalid payload`,
			http.StatusBadRequest,
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
		})
	}
}
