package application

import "truck-management/truck-management/domain"

type ILocationRepository interface {
	CreateLocation(truckID int, location *domain.Location) error
	GetLastLocation(truckID int) (*domain.Location, error)
}

type ITruckService interface {
	GetTruck(ID int) (*domain.Truck, error)
}

type ITripService interface {
	UpdateTrip(location domain.Location) error
}

type LocationService struct {
	locationRepository ILocationRepository
	truckService       ITruckService
	tripService        ITripService
}

func NewLocationService(repository ILocationRepository, truckService ITruckService, tripService ITripService) *LocationService {
	return &LocationService{
		locationRepository: repository,
		truckService:       truckService,
		tripService:        tripService,
	}
}

func (l *LocationService) CreateLocation(truckID int, location domain.Location) (domain.Location, error) {
	if _, err := l.truckService.GetTruck(truckID); err != nil {
		return location, err
	}

	if err := l.locationRepository.CreateLocation(truckID, &location); err != nil {
		return location, err
	}

	return location, l.tripService.UpdateTrip(location)
}

func (l *LocationService) GetLastLocation(truckID int) (*domain.Location, error) {
	if _, err := l.truckService.GetTruck(truckID); err != nil {
		return &domain.Location{}, err
	}

	return l.locationRepository.GetLastLocation(truckID)
}
