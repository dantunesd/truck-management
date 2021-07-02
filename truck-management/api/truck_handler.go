package api

import (
	"encoding/json"
	"net/http"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"

	"github.com/go-playground/validator"
)

func CreateTruckHandler(service *application.TruckService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var truck domain.Truck

		if dErr := json.NewDecoder(r.Body).Decode(&truck); dErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{"invalid content"})
			return
		}

		if vErr := validator.New().Struct(truck); vErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{vErr.Error()})
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

func GetTruckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseWriter(w, http.StatusOK, domain.Truck{})
	}
}
