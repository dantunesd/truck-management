package api

import (
	"fmt"
	"net/http"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"
)

func CreateTruckHandler(service *application.CreateTruckService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var truck domain.Truck

		result, err := service.CreateNewTruck(truck)
		fmt.Println(result, err)
	}
}
