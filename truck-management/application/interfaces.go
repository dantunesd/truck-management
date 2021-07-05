package application

import "truck-management/truck-management/domain"

type ITruckRepository interface {
	CreateTruck(truck *domain.Truck) error
	GetTruck(ID int) (domain.Truck, error)
	DeleteTruck(ID int) error
	UpdateTruck(ID int, truck *domain.Truck) error
}

type ILocationRepository interface {
	CreateLocation(truckID int, location *domain.Location) error
	GetLastLocation(truckID int) (domain.Location, error)
}

type ITripRepository interface {
	GetTrip(truckID int) (domain.Trip, error)
	UpsertTrip(trip *domain.Trip) error
}

type ITruckService interface {
	GetTruck(ID int) (domain.Truck, error)
}

type ITripService interface {
	UpdateTrip(location domain.Location) error
}

type ITripUpdater interface {
	UpdateTrip(currentTrip domain.Trip, location domain.Location) domain.Trip
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

type CreateLocationInput struct {
	EldID        string             `json:"eld_id" binding:"required,ascii,max=20"`
	EngineState  domain.EngineState `json:"engine_state" binding:"required,ascii,oneof=ON OFF"`
	CurrentSpeed int                `json:"current_speed" binding:"numeric"`
	Latitude     int                `json:"latitude" binding:"required,numeric"`
	Longitude    int                `json:"longitude" binding:"required,numeric"`
	EngineHours  int                `json:"engine_hours" binding:"required,numeric"`
	Odometer     int                `json:"odometer" binding:"required,numeric"`
}
