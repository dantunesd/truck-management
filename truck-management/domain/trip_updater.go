package domain

import "fmt"

type TripUpdater struct{}

func (t TripUpdater) UpdateTrip(currentTrip Trip, location Location) Trip {
	updatedTrip := Trip{
		ID:          currentTrip.ID,
		Destination: fmt.Sprintf("%d %d", location.Latitude, location.Longitude),
		State:       getState(location.EngineState),
	}

	if isANewTrip(currentTrip) {
		updatedTrip.TruckID = location.TruckID
		updatedTrip.Origin = fmt.Sprintf("%d %d", location.Latitude, location.Longitude)
		updatedTrip.Odometer = location.Odometer
		updatedTrip.EngineHours = location.EngineHours
		updatedTrip.AverageSpeed = location.Odometer / location.EngineHours

		return updatedTrip
	}

	updatedTrip.TruckID = currentTrip.TruckID
	updatedTrip.Origin = currentTrip.Origin
	updatedTrip.Odometer = currentTrip.Odometer + location.Odometer
	updatedTrip.EngineHours = currentTrip.EngineHours + location.EngineHours
	updatedTrip.AverageSpeed = updatedTrip.Odometer / updatedTrip.EngineHours

	return updatedTrip
}

func isANewTrip(trip Trip) bool {
	return trip.ID == 0 || trip.State == FINISHED
}

func getState(engine EngineState) TripState {
	if engine == ON {
		return ONGOING
	}
	return FINISHED
}
