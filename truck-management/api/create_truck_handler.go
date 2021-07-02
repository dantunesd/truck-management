package api

import (
	"encoding/json"
	"net/http"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"
)

func CreateTruckHandler(service *application.CreateTruckService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var truck domain.Truck

		if dErr := json.NewDecoder(r.Body).Decode(&truck); dErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{"invalid content"})
			return
		}

		result, err := service.CreateNewTruck(truck)
		if err != nil {
			responseWriter(w, getHTTPCode(err), &ErrorResponse{err.Error()})
			return
		}

		responseWriter(w, http.StatusOK, result)
	}
}
