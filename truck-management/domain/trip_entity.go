package domain

type TripState string

const (
	ONGOING  TripState = "ONGOING"
	FINISHED TripState = "FINISHED"
)

type Trip struct {
	ID           int    `json:"id"`
	EldID        string `json:"eld_id"`
	Origin       int    `json:"origin"`
	Destination  int    `json:"destination"`
	State        string `json:"state"`
	Odometer     string `json:"odometer"`
	EngineHours  string `json:"engine_hours"`
	AverageSpeed string `json:"average_speed"`
	UpdatedAt    string `json:"updated_at"`
}
