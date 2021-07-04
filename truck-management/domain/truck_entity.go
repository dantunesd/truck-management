package domain

type Truck struct {
	ID           int    `json:"id"`
	LicensePlate string `json:"license_plate" binding:"required,alphanum,min=7,max=7" gorm:"id"`
	EldID        string `json:"eld_id" binding:"required,ascii,max=20"`
	CarrierID    string `json:"carrier_id" binding:"required,ascii,max=50"`
	Type         string `json:"type" binding:"ascii,max=20"`
	Size         int    `json:"size" binding:"numeric,min=20,max=50"`
	Color        string `json:"color" binding:"ascii,max=20"`
	Make         string `json:"make" binding:"ascii,max=20"`
	Model        string `json:"model" binding:"ascii,max=20"`
	Year         int    `json:"year" binding:"numeric,min=1900"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
