package application

import "truck-management/truck-management/domain"

type ILocationRepository interface {
	CreateLocation(truckID int, location *domain.Location) error
	GetLastLocation(truckID int) (*domain.Location, error)
}

type LocationService struct {
	locationRepository ILocationRepository
}

func NewLocationService(repository ILocationRepository) *LocationService {
	return &LocationService{
		locationRepository: repository,
	}
}

func (l *LocationService) CreateLocation(truckID int, location domain.Location) (domain.Location, error) {
	return location, l.locationRepository.CreateLocation(truckID, &location)
}
