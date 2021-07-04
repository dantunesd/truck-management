package infrastructure

import (
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
	timeNow := NowFormated
	truck.CreatedAt = timeNow
	truck.UpdatedAt = timeNow

	return getError(t.db.Create(&truck))
}

func (t *TruckRepository) GetTruck(ID int) (*domain.Truck, error) {
	var truck domain.Truck
	return &truck, getError(t.db.Find(&truck, ID))
}

func (t *TruckRepository) DeleteTruck(ID int) error {
	var truck domain.Truck
	return t.db.Delete(&truck, ID).Error
}

func (t *TruckRepository) UpdateTruck(truck *domain.Truck) error {
	truck.UpdatedAt = NowFormated
	result := t.db.Model(&truck).Where("id = ?", truck.ID).Updates(truck)
	return getError(result)
}
