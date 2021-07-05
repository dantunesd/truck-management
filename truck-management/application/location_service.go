package application

import "truck-management/truck-management/domain"

type ILocationRepository interface {
	CreateLocation(truckID int, location *domain.Location) error
	GetLastLocation(truckID int) (domain.Location, error)
}

type ITruckService interface {
	GetTruck(ID int) (domain.Truck, error)
}

type ITripService interface {
	UpdateTrip(location domain.Location) error
}

type LocationService struct {
	locationRepository ILocationRepository
	truckService       ITruckService
	tripService        ITripService
}

func NewLocationService(locationRepository ILocationRepository, truckService ITruckService, tripService ITripService) *LocationService {
	return &LocationService{
		locationRepository: locationRepository,
		truckService:       truckService,
		tripService:        tripService,
	}
}

type CreateLocationInput struct {
	EldID        string             `json:"eld_id" binding:"required,ascii,max=20"`
	EngineState  domain.EngineState `json:"engine_state" binding:"required,ascii,oneof=ON OFF"`
	CurrentSpeed int                `json:"current_speed" binding:"required,numeric,min=0,max=500"`
	Latitude     int                `json:"latitude" binding:"required,numeric,min=0"`
	Longitude    int                `json:"longitude" binding:"required,numeric,min=0"`
	EngineHours  int                `json:"engine_hours" binding:"required,numeric,min=0"`
	Odometer     int                `json:"odometer" binding:"required,numeric,min=0"`
}

func (l *LocationService) CreateLocation(truckID int, input CreateLocationInput) (domain.Location, error) {
	location := domain.Location{
		EldID:        input.EldID,
		EngineState:  input.EngineState,
		CurrentSpeed: input.CurrentSpeed,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		EngineHours:  input.EngineHours,
		Odometer:     input.Odometer,
	}

	if _, err := l.truckService.GetTruck(truckID); err != nil {
		return location, err
	}

	if err := l.locationRepository.CreateLocation(truckID, &location); err != nil {
		return location, err
	}

	return location, l.tripService.UpdateTrip(location)
}

func (l *LocationService) GetLastLocation(truckID int) (domain.Location, error) {
	if _, err := l.truckService.GetTruck(truckID); err != nil {
		return domain.Location{}, err
	}

	return l.locationRepository.GetLastLocation(truckID)
}
