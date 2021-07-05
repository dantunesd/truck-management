package domain

type TripUpdater struct{}

func (t TripUpdater) UpdateTrip(currentTrip Trip, location Location) Trip {
	if currentTrip.IsNewTrip() {
		return Trip{
			ID:           currentTrip.ID,
			Destination:  location.GetLatitudeAndLongitude(),
			State:        location.GetTripState(),
			TruckID:      location.TruckID,
			Origin:       location.GetLatitudeAndLongitude(),
			Odometer:     location.Odometer,
			EngineHours:  location.EngineHours,
			AverageSpeed: location.GetAverageSpeed(),
		}
	}

	return Trip{
		ID:           currentTrip.ID,
		Destination:  location.GetLatitudeAndLongitude(),
		State:        location.GetTripState(),
		TruckID:      currentTrip.TruckID,
		Origin:       currentTrip.Origin,
		Odometer:     currentTrip.Odometer + location.Odometer,
		EngineHours:  currentTrip.EngineHours + location.EngineHours,
		AverageSpeed: (currentTrip.Odometer + location.Odometer) / (currentTrip.EngineHours + location.EngineHours),
	}
}
