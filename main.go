package main

import (
	"fmt"
	"truck-management/truck-management/domain"
	"truck-management/truck-management/infrastructure"
)

func main() {
	createTruckService := infrastructure.CreateTruckServiceFactory()

	fromRequest := domain.Truck{
		LicensePlate: "ABC1234",
	}

	created, err := createTruckService.CreateNewTruck(fromRequest)

	fmt.Println(err)
	fmt.Printf("%+v", created)
}
