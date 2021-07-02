package application

import (
	"truck-management/truck-management/domain"
)

type ITruckRepository interface {
	CreateTruck(truck *domain.Truck) error
}

type TruckService struct {
	truckRepository ITruckRepository
}

func NewTruckService(repository ITruckRepository) *TruckService {
	return &TruckService{
		truckRepository: repository,
	}
}

func (c *TruckService) CreateNewTruck(newTruck domain.Truck) (domain.Truck, error) {
	return newTruck, c.truckRepository.CreateTruck(&newTruck)
}
