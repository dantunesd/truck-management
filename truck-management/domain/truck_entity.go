package domain

type Truck struct {
	LicensePlate string `json:"license_plate"`
	EldID        string `json:"eld_id"`
	Carrier      string `json:"carrier"`
	Size         int    `json:"size"`
	Color        string `json:"color"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         string `json:"year"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
