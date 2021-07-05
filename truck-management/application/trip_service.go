package application

import "truck-management/truck-management/domain"

type ITripRepository interface {
	GetTrip(truckID int) (*domain.Trip, error)
}

type TripService struct {
	tripRepository ITripRepository
	truckService   ITruckService
}

func NewTripService(repository ITripRepository, truckService ITruckService) *TripService {
	return &TripService{
		tripRepository: repository,
		truckService:   truckService,
	}
}

func (t *TripService) GetTrip(truckID int) (*domain.Trip, error) {
	if _, err := t.truckService.GetTruck(truckID); err != nil {
		return &domain.Trip{}, err
	}

	return t.tripRepository.GetTrip(truckID)
}
