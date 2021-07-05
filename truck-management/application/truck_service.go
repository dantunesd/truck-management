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

type TruckCreateInput struct {
	LicensePlate string `json:"license_plate" binding:"required,alphanum,min=7,max=7"`
	EldID        string `json:"eld_id" binding:"required,ascii,max=20"`
	CarrierID    string `json:"carrier_id" binding:"required,ascii,max=50"`
	Type         string `json:"type" binding:"ascii,max=20"`
	Size         int    `json:"size" binding:"numeric"`
	Color        string `json:"color" binding:"ascii,max=20"`
	Make         string `json:"make" binding:"ascii,max=20"`
	Model        string `json:"model" binding:"ascii,max=20"`
	Year         int    `json:"year" binding:"numeric"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
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

func (t *TruckService) GetTruck(ID int) (domain.Truck, error) {
	return t.truckRepository.GetTruck(ID)
}

type TruckUpdateInput struct {
	LicensePlate string `json:"license_plate" binding:"ascii,max=7"`
	EldID        string `json:"eld_id" binding:"ascii,max=20"`
	CarrierID    string `json:"carrier_id" binding:"ascii,max=50"`
	Type         string `json:"type" binding:"ascii,max=20"`
	Size         int    `json:"size" binding:"numeric"`
	Color        string `json:"color" binding:"ascii,max=20"`
	Make         string `json:"make" binding:"ascii,max=20"`
	Model        string `json:"model" binding:"ascii,max=20"`
	Year         int    `json:"year" binding:"numeric"`
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

func (t *TruckService) DeleteTruck(ID int) error {
	return t.truckRepository.DeleteTruck(ID)
}
