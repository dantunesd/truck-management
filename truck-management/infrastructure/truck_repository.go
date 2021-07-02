package infrastructure

import (
	"strings"
	"time"
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
		return domain.NewConflict("this truck is already registered")
	}

	return result.Error
}
