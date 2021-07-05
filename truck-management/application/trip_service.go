package application

import (
	"strings"
	"truck-management/truck-management/domain"
)

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

func (t *TripService) GetTrip(truckID int) (domain.Trip, error) {
	if _, err := t.truckService.GetTruck(truckID); err != nil {
		return domain.Trip{}, err
	}

	return t.tripRepository.GetTrip(truckID)
}

func (t *TripService) UpdateTrip(location domain.Location) error {
	currentTrip, err := t.tripRepository.GetTrip(location.TruckID)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return err
	}

	updatedTrip := t.tripUpdater.UpdateTrip(currentTrip, location)

	return t.tripRepository.UpsertTrip(&updatedTrip)
}
