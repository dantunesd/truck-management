package domain

type Truck struct {
	ID           int    `json:"id"`
	LicensePlate string `json:"license_plate" validate:"required,alphanum,min=7,max=7"`
	EldID        string `json:"eld_id" validate:"required,ascii,max=20"`
	CarrierID    string `json:"carrier_id" validate:"required,ascii,max=50"`
	Type         string `json:"type" validate:"ascii,max=20"`
	Size         int    `json:"size" validate:"numeric,min=20,max=50"`
	Color        string `json:"color" validate:"ascii,max=20"`
	Make         string `json:"make" validate:"ascii,max=20"`
	Model        string `json:"model" validate:"ascii,max=20"`
	Year         int    `json:"year" validate:"numeric,min=1900"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
