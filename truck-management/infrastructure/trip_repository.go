package infrastructure

import (
	"time"
	"truck-management/truck-management/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{
		db: db,
	}
}

func (t *TripRepository) GetTrip(truckID int) (domain.Trip, error) {
	var trip domain.Trip

	result := t.db.Where("truck_id = ?", truckID).Find(&trip)

	if result.Error != nil {
		return trip, result.Error
	}

	if isNotFound(result) {
		return trip, TripNotFoundError
	}

	return trip, nil
}

func (t *TripRepository) UpsertTrip(trip *domain.Trip) error {
	trip.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return t.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&trip).Error
}
