package domain

type TripState string

const (
	ONGOING  TripState = "ONGOING"
	FINISHED TripState = "FINISHED"
)

type Trip struct {
	ID           int       `json:"id,omitempty"`
	TruckID      int       `json:"truck_id,omitempty"`
	Origin       string    `json:"origin,omitempty"`
	Destination  string    `json:"destination,omitempty"`
	State        TripState `json:"state,omitempty"`
	Odometer     int       `json:"odometer,omitempty"`
	EngineHours  int       `json:"engine_hours,omitempty"`
	AverageSpeed int       `json:"average_speed,omitempty"`
	UpdatedAt    string    `json:"updated_at,omitempty"`
}

func (t Trip) IsNewTrip() bool {
	return t.ID == 0 || t.State == FINISHED
}
