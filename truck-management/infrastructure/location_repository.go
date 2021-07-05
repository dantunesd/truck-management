package infrastructure

import (
	"time"
	"truck-management/truck-management/domain"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func (t *LocationRepository) CreateLocation(truckID int, location *domain.Location) error {
	location.TruckID = truckID
	location.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	return t.db.Create(&location).Error
}

func (t *LocationRepository) GetLastLocation(truckID int) (domain.Location, error) {
	var location domain.Location

	return location, t.db.Where("truck_id = ?", truckID).Order("id DESC").Limit(1).Find(&location).Error
}
