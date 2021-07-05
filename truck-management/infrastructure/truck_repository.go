package infrastructure

import (
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

	if isDuplicated(result) {
		return ConflictError
	}

	return result.Error

}

func (t *TruckRepository) GetTruck(ID int) (domain.Truck, error) {
	var truck domain.Truck

	result := t.db.Find(&truck, ID)

	if result.Error != nil {
		return truck, result.Error
	}

	if isNotFound(result) {
		return truck, NotFoundError
	}

	return truck, nil
}

func (t *TruckRepository) DeleteTruck(ID int) error {
	var truck domain.Truck

	return t.db.Delete(&truck, ID).Error
}

func (t *TruckRepository) UpdateTruck(ID int, truck *domain.Truck) error {
	truck.ID = ID
	truck.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	result := t.db.Model(&truck).Where("id = ?", ID).Updates(truck)

	if isDuplicated(result) {
		return ConflictError
	}

	return result.Error
}
