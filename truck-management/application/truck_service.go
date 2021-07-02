package application

import (
	"truck-management/truck-management/domain"
)

type ITruckRepository interface {
	CreateTruck(truck *domain.Truck) error
}

type ITruckValidator interface {
	IsValidTruck(truck domain.Truck) error
}

type TruckService struct {
	TruckRepository ITruckRepository
	TruckValidator  ITruckValidator
}

func (c *TruckService) CreateNewTruck(newTruck domain.Truck) (domain.Truck, error) {
	if tErr := c.TruckValidator.IsValidTruck(newTruck); tErr != nil {
		return newTruck, tErr
	}

	if cErr := c.TruckRepository.CreateTruck(&newTruck); cErr != nil {
		return newTruck, cErr
	}

	return newTruck, nil
}
