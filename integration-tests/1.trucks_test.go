package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestTrucks(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		method      string
		body        string
		wantStatus  int
		wantReponse string
	}{
		{
			"Should return 201 when creating a truck with success",
			"http://app:3000/trucks",
			http.MethodPost,
			`{"license_plate":"ABC1234","eld_id":"eld-id","carrier_id":"My Carrier","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900}`,
			http.StatusCreated,
			``,
		},
		{
			"Should return 409 when creating a truck with the same license plate",
			"http://app:3000/trucks",
			http.MethodPost,
			`{"license_plate":"ABC1234","eld_id":"new-eld-id","carrier_id":"My Carrier","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900}`,
			http.StatusConflict,
			`{"title":"license plate or eld_id is already registered","status":409}`,
		},
		{
			"Should return 409 when creating a truck with the same eld id",
			"http://app:3000/trucks",
			http.MethodPost,
			`{"license_plate":"ABC4567","eld_id":"eld-id","carrier_id":"My Carrier","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900}`,
			http.StatusConflict,
			`{"title":"license plate or eld_id is already registered","status":409}`,
		},
		{
			"Should return 400 when creating a truck without the required fields",
			"http://app:3000/trucks",
			http.MethodPost,
			`{}`,
			http.StatusBadRequest,
			`{"title":"Key: 'TruckCreateInput.LicensePlate' Error:Field validation for 'LicensePlate' failed on the 'required' tag\nKey: 'TruckCreateInput.EldID' Error:Field validation for 'EldID' failed on the 'required' tag\nKey: 'TruckCreateInput.CarrierID' Error:Field validation for 'CarrierID' failed on the 'required' tag","status":400}`,
		},
		{
			"Should return 200 when getting an existing truck",
			"http://app:3000/trucks/1",
			http.MethodGet,
			``,
			http.StatusOK,
			`{"id":1,"license_plate":"INSERTT","eld_id":"INSERTED ELD","carrier_id":"INSERTT MY CARRIER","type":"REEFERR","size":23,"color":"blue","make":"Maker","model":"Model","year":1900,"created_at":"2021-07-04 21:34:26","updated_at":"2021-07-04 21:34:26"}`,
		},
		{
			"Should return 404 when getting an inexistent truck",
			"http://app:3000/trucks/9999",
			http.MethodGet,
			``,
			http.StatusNotFound,
			`{"title":"truck not found","status":404}`,
		},
		{
			"Should return 400 when getting a truck with an invalid truck id",
			"http://app:3000/trucks/invalid",
			http.MethodGet,
			``,
			http.StatusBadRequest,
			`{"title":"id must be numeric","status":400}`,
		},
		{
			"Should return 204 when updating an existing truck",
			"http://app:3000/trucks/1",
			http.MethodPatch,
			`{"license_plate":"NEWPLAT","eld_id":"NEW-ELD-ID","carrier_id":"NEW My Carrier","type":"NEW REEFERR"}`,
			http.StatusNoContent,
			``,
		},
		{
			"Should return 404 when updating an inexistent truck",
			"http://app:3000/trucks/9999",
			http.MethodPatch,
			`{"license_plate":"NEWPLAT","eld_id":"NEW-ELD-ID","carrier_id":"NEW My Carrier","type":"NEW REEFERR"}`,
			http.StatusNotFound,
			`{"title":"truck not found","status":404}`,
		},
		{
			"Should return 404 when updating a truck with an invalid truck id",
			"http://app:3000/trucks/invalid",
			http.MethodPatch,
			`{"license_plate":"NEWPLAT","eld_id":"NEW-ELD-ID","carrier_id":"NEW My Carrier","type":"NEW REEFERR"}`,
			http.StatusBadRequest,
			`{"title":"id must be numeric","status":400}`,
		},
		{
			"Should return 204 when updating a truck without the required fields",
			"http://app:3000/trucks/1",
			http.MethodPatch,
			`{}`,
			http.StatusNoContent,
			``,
		},
		{
			"Should return 204 when deleting an existing truck",
			"http://app:3000/trucks/2",
			http.MethodDelete,
			``,
			http.StatusNoContent,
			``,
		},
		{
			"Should return 204 when deleting an inexistent truck",
			"http://app:3000/trucks/9999",
			http.MethodDelete,
			``,
			http.StatusNoContent,
			``,
		},
		{
			"Should return 400 when deleting a truck with an invalid truck id",
			"http://app:3000/trucks/invalid",
			http.MethodDelete,
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
