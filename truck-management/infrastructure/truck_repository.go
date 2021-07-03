package infrastructure

import (
	"strings"
	"time"
	"truck-management/truck-management/api"
	"truck-management/truck-management/domain"

	"gorm.io/gorm"
)

type TruckRepository struct {
	db *gorm.DB
}

func NewTruckRepository(db *gorm.DB) *TruckRepository {
	return &TruckRepository{
		db: db,
	}
}

func (t *TruckRepository) CreateTruck(truck *domain.Truck) error {
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	truck.CreatedAt = timeNow
	truck.UpdatedAt = timeNow

	result := t.db.Create(&truck)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return api.NewConflict("license plate or eld_id is already registered")
	}

	return result.Error
}

func (t *TruckRepository) GetTruck(ID int) (*domain.Truck, error) {
	var truck domain.Truck
	result := t.db.Find(&truck, ID)

	if result.RowsAffected == 0 {
		return &truck, api.NewNotFound("truck not found")
	}

	return &truck, result.Error
}

func (t *TruckRepository) DeleteTruck(ID int) error {
	var truck domain.Truck
	return t.db.Delete(&truck, ID).Error
}
