package application

import "truck-management/truck-management/domain"

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
