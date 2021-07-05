package domain

type TripUpdater struct{}

func NewTripUpdater() *TripUpdater {
	return &TripUpdater{}
}

func (t *TripUpdater) UpdateTrip(currentTrip Trip, location Location) Trip {
	if currentTrip.IsNewTrip() {
		return Trip{
			ID:           currentTrip.ID,
			TruckID:      location.TruckID,
			Origin:       location.GetLatitudeAndLongitude(),
			Destination:  location.GetLatitudeAndLongitude(),
			State:        location.GetTripState(),
			Odometer:     location.Odometer,
			EngineHours:  location.EngineHours,
			AverageSpeed: location.GetAverageSpeed(),
		}
	}

	return Trip{
		ID:           currentTrip.ID,
		TruckID:      currentTrip.TruckID,
		Origin:       currentTrip.Origin,
		Destination:  location.GetLatitudeAndLongitude(),
		State:        location.GetTripState(),
		Odometer:     currentTrip.Odometer + location.Odometer,
		EngineHours:  currentTrip.EngineHours + location.EngineHours,
		AverageSpeed: (currentTrip.Odometer + location.Odometer) / (currentTrip.EngineHours + location.EngineHours),
	}
}
