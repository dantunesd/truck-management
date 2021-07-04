package application

import "truck-management/truck-management/domain"

type ILocationRepository interface {
	CreateLocation(truckID int, location *domain.Location) error
	GetLastLocation(truckID int) (*domain.Location, error)
}

type ITruckService interface {
	GetTruck(ID int) (*domain.Truck, error)
}

type LocationService struct {
	locationRepository ILocationRepository
	truckService       ITruckService
}

func NewLocationService(repository ILocationRepository, truckService ITruckService) *LocationService {
	return &LocationService{
		locationRepository: repository,
		truckService:       truckService,
	}
}

func (l *LocationService) CreateLocation(truckID int, location domain.Location) (domain.Location, error) {
	if _, err := l.truckService.GetTruck(truckID); err != nil {
		return location, err
	}

	return location, l.locationRepository.CreateLocation(truckID, &location)
}
