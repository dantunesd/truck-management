package domain

type EngineState string

const (
	ON  EngineState = "ON"
	OFF EngineState = "OFF"
)

type Location struct {
	ID           int         `json:"id"`
	TruckID      string      `json:"truck_id"`
	EldID        string      `json:"eld_id"`
	EngineState  EngineState `json:"engine_state"`
	CurrentSpeed float64     `json:"current_speed"`
	Latitude     int         `json:"latitude"`
	Longitude    int         `json:"longitude"`
	EngineHours  string      `json:"engine_hours"`
	Odometer     string      `json:"odometer"`
	CreatedAt    string      `json:"created_at"`
}
