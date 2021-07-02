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
	truckRepository ITruckRepository
	truckValidator  ITruckValidator
}

func NewTruckService(repository ITruckRepository, validator ITruckValidator) *TruckService {
	return &TruckService{
		truckRepository: repository,
		truckValidator:  validator,
	}
}

func (c *TruckService) CreateNewTruck(newTruck domain.Truck) (domain.Truck, error) {
	if tErr := c.truckValidator.IsValidTruck(newTruck); tErr != nil {
		return newTruck, tErr
	}

	if cErr := c.truckRepository.CreateTruck(&newTruck); cErr != nil {
		return newTruck, cErr
	}

	return newTruck, nil
}
