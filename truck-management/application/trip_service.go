package application

import "truck-management/truck-management/domain"

type ITripRepository interface {
	GetTrip(truckID int) (*domain.Trip, error)
	SaveTrip(trip *domain.Trip) error
}

type ITripUpdater interface {
	UpdateTrip(currentTrip domain.Trip, location domain.Location) domain.Trip
}

type TripService struct {
	tripRepository ITripRepository
	truckService   ITruckService
	tripUpdater    ITripUpdater
}

func NewTripService(tripRepository ITripRepository, truckService ITruckService, tripUpdater ITripUpdater) *TripService {
	return &TripService{
		tripRepository: tripRepository,
		truckService:   truckService,
		tripUpdater:    tripUpdater,
	}
}

func (t *TripService) GetTrip(truckID int) (*domain.Trip, error) {
	if _, err := t.truckService.GetTruck(truckID); err != nil {
		return &domain.Trip{}, err
	}

	return t.tripRepository.GetTrip(truckID)
}

func (t *TripService) UpdateTrip(location domain.Location) error {
	currentTrip, err := t.tripRepository.GetTrip(location.TruckID)
	if err != nil {
		return err
	}

	updatedTrip := t.tripUpdater.UpdateTrip(*currentTrip, location)

	return t.tripRepository.SaveTrip(&updatedTrip)
}
