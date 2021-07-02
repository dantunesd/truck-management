package domain

type Truck struct {
	ID           string `json:"id"`
	LicensePlate string `json:"license_plate" validate:"required,alphanum,min=7,max=7"`
	EldID        string `json:"eld_id" validate:"required,ascii,max=20"`
	Carrier      string `json:"carrier" validate:"required,ascii,max=50"`
	Type         string `json:"type"`
	Size         int    `json:"size"`
	Color        string `json:"color"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
