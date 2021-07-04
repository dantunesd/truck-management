package domain

type EngineState string

const (
	ON  EngineState = "ON"
	OFF EngineState = "OFF"
)

type Location struct {
	ID           int         `json:"id"`
	TruckID      int         `json:"truck_id"`
	EldID        string      `json:"eld_id" binding:"required,ascii,max=20"`
	EngineState  EngineState `json:"engine_state" binding:"required,ascii,oneof=ON OFF"`
	CurrentSpeed int         `json:"current_speed" binding:"required,numeric,min=0,max=500"` // in KM Ex: 95km/h, 120km/h
	Latitude     int         `json:"latitude" binding:"required,numeric,min=0"`
	Longitude    int         `json:"longitude" binding:"required,numeric,min=0"`
	EngineHours  int         `json:"engine_hours" binding:"required,numeric,min=0"` // in hours Ex: 1h, 3h
	Odometer     int         `json:"odometer" binding:"required,numeric,min=0"`     // in KM	Ex: 5km, 70km
	CreatedAt    string      `json:"created_at"`
}
