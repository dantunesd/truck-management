package domain

type TripState string

const (
	ONGOING  TripState = "ONGOING"
	FINISHED TripState = "FINISHED"
)

type Trip struct {
	ID           int       `json:"id"`
	TruckID      int       `json:"truck_id"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	State        TripState `json:"state"`
	Odometer     int       `json:"odometer"`
	EngineHours  int       `json:"engine_hours"`
	AverageSpeed int       `json:"average_speed"`
	UpdatedAt    string    `json:"updated_at"`
}

func (t *Trip) IsNewTrip() bool {
	return t.ID == 0 || t.State == FINISHED
}
