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
	if result.Error != nil && isDuplicated(result.Error) {
		return api.NewConflict("license plate or eld_id is already registered")
	}

	return result.Error
}

func (t *TruckRepository) GetTruck(ID int) (*domain.Truck, error) {
	var truck domain.Truck
	result := t.db.Find(&truck, ID)

	if err := getError(result); err != nil {
		return &truck, err
	}

	return &truck, nil
}

func (t *TruckRepository) DeleteTruck(ID int) error {
	var truck domain.Truck
	return t.db.Delete(&truck, ID).Error
}

func (t *TruckRepository) UpdateTruck(truck *domain.Truck) error {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	truck.UpdatedAt = timeNow

	result := t.db.Model(&truck).Where("id = ?", truck.ID).Updates(truck)

	return getError(result)
}

func getError(result *gorm.DB) error {
	if result.Error != nil {
		if isDuplicated(result.Error) {
			return api.NewConflict("license plate or eld_id is already registered")
		}
		return result.Error
	}

	if result.RowsAffected == 0 {
		return api.NewNotFound("truck not found")
	}

	return nil
}

func isDuplicated(err error) bool {
	return strings.Contains(err.Error(), "Duplicate entry")
}
