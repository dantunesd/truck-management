package domain

import "fmt"

type EngineState string

const (
	ON  EngineState = "ON"
	OFF EngineState = "OFF"
)

type Location struct {
	ID           int         `json:"id,omitempty"`
	TruckID      int         `json:"truck_id,omitempty"`
	EldID        string      `json:"eld_id,omitempty"`
	EngineState  EngineState `json:"engine_state,omitempty"`
	CurrentSpeed int         `json:"current_speed,omitempty"`
	Latitude     int         `json:"latitude,omitempty"`
	Longitude    int         `json:"longitude,omitempty"`
	EngineHours  int         `json:"engine_hours,omitempty"`
	Odometer     int         `json:"odometer,omitempty"`
	CreatedAt    string      `json:"created_at,omitempty"`
}

func (l Location) GetLatitudeAndLongitude() string {
	return fmt.Sprintf("%d %d", l.Latitude, l.Longitude)
}

func (l Location) GetAverageSpeed() int {
	return l.Odometer / l.EngineHours
}

func (l Location) GetTripState() TripState {
	if l.EngineState == ON {
		return ONGOING
	}

	return FINISHED
}
