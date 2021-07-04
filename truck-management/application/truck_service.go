package application

import (
	"truck-management/truck-management/domain"
)

type ITruckRepository interface {
	CreateTruck(truck *domain.Truck) error
	GetTruck(ID int) (*domain.Truck, error)
	DeleteTruck(ID int) error
	UpdateTruck(ID int, truck *domain.Truck) error
}

type TruckService struct {
	truckRepository ITruckRepository
}

func NewTruckService(repository ITruckRepository) *TruckService {
	return &TruckService{
		truckRepository: repository,
	}
}

func (c *TruckService) CreateTruck(newTruck domain.Truck) (domain.Truck, error) {
	return newTruck, c.truckRepository.CreateTruck(&newTruck)
}

func (c *TruckService) GetTruck(ID int) (*domain.Truck, error) {
	return c.truckRepository.GetTruck(ID)
}

func (c *TruckService) UpdateTruck(ID int, truck domain.Truck) error {
	return c.truckRepository.UpdateTruck(ID, &truck)
}

func (c *TruckService) DeleteTruck(ID int) error {
	return c.truckRepository.DeleteTruck(ID)
}
