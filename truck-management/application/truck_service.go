package application

import (
	"truck-management/truck-management/domain"
)

type TruckService struct {
	truckRepository ITruckRepository
}

func NewTruckService(truckRepository ITruckRepository) *TruckService {
	return &TruckService{
		truckRepository: truckRepository,
	}
}

func (t *TruckService) CreateTruck(input TruckCreateInput) (domain.Truck, error) {
	truck := domain.Truck{
		LicensePlate: input.LicensePlate,
		EldID:        input.EldID,
		CarrierID:    input.CarrierID,
		Type:         input.Type,
		Size:         input.Size,
		Color:        input.Color,
		Make:         input.Make,
		Model:        input.Model,
		Year:         input.Year,
	}

	return truck, t.truckRepository.CreateTruck(&truck)
}

func (t *TruckService) UpdateTruck(ID int, input TruckUpdateInput) error {
	truck := domain.Truck{
		LicensePlate: input.LicensePlate,
		EldID:        input.EldID,
		CarrierID:    input.CarrierID,
		Type:         input.Type,
		Size:         input.Size,
		Color:        input.Color,
		Make:         input.Make,
		Model:        input.Model,
		Year:         input.Year,
	}

	if _, err := t.GetTruck(ID); err != nil {
		return err
	}

	return t.truckRepository.UpdateTruck(ID, &truck)
}

func (t *TruckService) GetTruck(ID int) (domain.Truck, error) {
	return t.truckRepository.GetTruck(ID)
}

func (t *TruckService) DeleteTruck(ID int) error {
	return t.truckRepository.DeleteTruck(ID)
}
