package infrastructure

import (
	"strings"
	"time"
	"truck-management/truck-management/api"

	"gorm.io/gorm"
)

var NowFormated = time.Now().Format("2006-01-02 15:04:05")
var NotFoundError = api.NewNotFound("truck not found")
var ConflictError = api.NewConflict("license plate or eld_id is already registered")

func hasError(result *gorm.DB) bool {
	return result.Error != nil
}

func isDuplicated(result *gorm.DB) bool {
	return hasError(result) && strings.Contains(result.Error.Error(), "Duplicate entry")
}

func isNotFound(result *gorm.DB) bool {
	return result.RowsAffected == 0
}

func getError(result *gorm.DB) error {
	if isDuplicated(result) {
		return ConflictError
	}

	if hasError(result) {
		return result.Error
	}

	if isNotFound(result) {
		return NotFoundError
	}

	return nil
}
