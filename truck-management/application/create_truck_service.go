package application

import (
	"truck-management/truck-management/domain"
)

type TruckRepository interface {
	CreateTruck(truck *domain.Truck) error
}

type TruckValidator interface {
	IsValidTruck(newTruck domain.Truck) error
}

type CreateTruckService struct {
	TruckRepository TruckRepository
	TruckValidator  TruckValidator
}

func (c *CreateTruckService) CreateNewTruck(newTruck domain.Truck) (domain.Truck, error) {
	if tErr := c.TruckValidator.IsValidTruck(newTruck); tErr != nil {
		return newTruck, tErr
	}

	if cErr := c.TruckRepository.CreateTruck(&newTruck); cErr != nil {
		return newTruck, cErr
	}

	return newTruck, nil
}
