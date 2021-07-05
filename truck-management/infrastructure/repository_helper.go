package infrastructure

import (
	"strings"
	"truck-management/truck-management/api"

	"gorm.io/gorm"
)

var NotFoundError = api.NewNotFound("truck not found")
var ConflictError = api.NewConflict("license plate or eld_id is already registered")

func isDuplicated(result *gorm.DB) bool {
	return result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry")
}

func isNotFound(result *gorm.DB) bool {
	return result.RowsAffected == 0
}
