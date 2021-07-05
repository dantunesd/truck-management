package infrastructure

import (
	"truck-management/truck-management/domain"

	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{
		db: db,
	}
}

func (t *TripRepository) GetTrip(truckID int) (*domain.Trip, error) {
	var trip domain.Trip
	result := t.db.Where("truck_id = ?", truckID).Find(&trip)
	return &trip, result.Error
}

func (t *TripRepository) SaveTrip(trip *domain.Trip) error {
	return nil
}
