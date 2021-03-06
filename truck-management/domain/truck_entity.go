package domain

type Truck struct {
	ID           int    `json:"id"`
	LicensePlate string `json:"license_plate"`
	EldID        string `json:"eld_id"`
	CarrierID    string `json:"carrier_id"`
	Type         string `json:"type"`
	Size         int    `json:"size"`
	Color        string `json:"color"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
