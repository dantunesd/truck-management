package infrastructure

import (
	"time"
	"truck-management/truck-management/domain"
)

type TruckRepository struct {
}

func NewTruckRepository() *TruckRepository {
	return &TruckRepository{}
}

func (t *TruckRepository) CreateTruck(truck *domain.Truck) error {
	timeNow := time.Now().Format(time.RFC3339)

	truck.ID = "NEWID"
	truck.CreatedAt = timeNow
	truck.UpdatedAt = timeNow

	// database implementation goes here
	return nil
}
