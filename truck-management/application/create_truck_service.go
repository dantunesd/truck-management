package application

import (
	"truck-management/truck-management/domain"
)

type TruckRepository interface {
	GetTruckByLicensePlateAndEldID(licensePlate, eldID string) (domain.Truck, error)
	CreateTruck(truck *domain.Truck) error
}

type TruckValidator func(newTruck domain.Truck, possibleExistingTruck domain.Truck) error

type CreateTruckService struct {
	TruckRepository TruckRepository
	TruckValidator  TruckValidator
}

func (c *CreateTruckService) CreateNewTruck(newTruck domain.Truck) (domain.Truck, error) {
	possibleExistingTruck, gErr := c.TruckRepository.GetTruckByLicensePlateAndEldID(newTruck.LicensePlate, newTruck.EldID)

	if gErr != nil {
		return newTruck, gErr
	}

	if tErr := c.TruckValidator(newTruck, possibleExistingTruck); tErr != nil {
		return newTruck, tErr
	}

	if cErr := c.TruckRepository.CreateTruck(&newTruck); cErr != nil {
		return newTruck, cErr
	}

	return newTruck, nil
}
