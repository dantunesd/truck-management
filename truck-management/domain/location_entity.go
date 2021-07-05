package domain

import "fmt"

type EngineState string

const (
	ON  EngineState = "ON"
	OFF EngineState = "OFF"
)

type Location struct {
	ID           int         `json:"id"`
	TruckID      int         `json:"truck_id"`
	EldID        string      `json:"eld_id"`
	EngineState  EngineState `json:"engine_state"`
	CurrentSpeed int         `json:"current_speed"`
	Latitude     int         `json:"latitude"`
	Longitude    int         `json:"longitude"`
	EngineHours  int         `json:"engine_hours"`
	Odometer     int         `json:"odometer"`
	CreatedAt    string      `json:"created_at"`
}

func (l *Location) GetLatitudeAndLongitude() string {
	return fmt.Sprintf("%d %d", l.Latitude, l.Longitude)
}

func (l *Location) GetAverageSpeed() int {
	return l.Odometer / l.EngineHours
}

func (l *Location) GetTripState() TripState {
	if l.EngineState == ON {
		return ONGOING
	}

	return FINISHED
}
