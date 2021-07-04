package domain

type EngineState string

const (
	ON  EngineState = "ON"
	OFF EngineState = "OFF"
)

type Location struct {
	ID           int         `json:"id,omitempty"`
	TruckID      int         `json:"truck_id,omitempty"`
	EldID        string      `json:"eld_id,omitempty" binding:"required,ascii,max=20"`
	EngineState  EngineState `json:"engine_state,omitempty" binding:"required,ascii,oneof=ON OFF"`
	CurrentSpeed int         `json:"current_speed,omitempty" binding:"required,numeric,min=0,max=500"` // in KM Ex: 95km/h, 120km/h
	Latitude     int         `json:"latitude,omitempty" binding:"required,numeric,min=0"`
	Longitude    int         `json:"longitude,omitempty" binding:"required,numeric,min=0"`
	EngineHours  int         `json:"engine_hours,omitempty" binding:"required,numeric,min=0"` // in hours Ex: 1h, 3h
	Odometer     int         `json:"odometer,omitempty" binding:"required,numeric,min=0"`     // in KM	Ex: 5km, 70km
	CreatedAt    string      `json:"created_at,omitempty"`
}
