package domain

type TripState string

const (
	ONGOING  TripState = "ONGOING"
	FINISHED TripState = "FINISHED"
)

type Trip struct {
	ID           int       `json:"id,omitempty"`
	TruckID      int       `json:"truck_id,omitempty"`
	Origin       int       `json:"origin,omitempty"`
	Destination  int       `json:"destination,omitempty"`
	State        TripState `json:"state,omitempty"`
	Odometer     int       `json:"odometer,omitempty"`
	EngineHours  int       `json:"engine_hours,omitempty"`
	AverageSpeed int       `json:"average_speed,omitempty"`
	UpdatedAt    string    `json:"updated_at,omitempty"`
}
