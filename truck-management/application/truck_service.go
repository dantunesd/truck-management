package application

import (
	"truck-management/truck-management/domain"
)

type ITruckRepository interface {
	CreateTruck(truck *domain.Truck) error
	GetTruck(ID int) (domain.Truck, error)
	DeleteTruck(ID int) error
	UpdateTruck(ID int, truck *domain.Truck) error
}

type TruckService struct {
	truckRepository ITruckRepository
}

func NewTruckService(truckRepository ITruckRepository) *TruckService {
	return &TruckService{
		truckRepository: truckRepository,
	}
}

func (t *TruckService) CreateTruck(truck domain.Truck) (domain.Truck, error) {
	return truck, t.truckRepository.CreateTruck(&truck)
}

func (t *TruckService) GetTruck(ID int) (domain.Truck, error) {
	return t.truckRepository.GetTruck(ID)
}

func (t *TruckService) UpdateTruck(ID int, truck domain.Truck) error {
	return t.truckRepository.UpdateTruck(ID, &truck)
}

func (t *TruckService) DeleteTruck(ID int) error {
	return t.truckRepository.DeleteTruck(ID)
}
