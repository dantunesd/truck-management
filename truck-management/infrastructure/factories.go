package infrastructure

import (
	"net/http"
	"truck-management/truck-management/api"
	"truck-management/truck-management/application"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseFactory(config *Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.DBConn), &gorm.Config{})
}

func TruckServiceFactory(truckRepository application.ITruckRepository) *application.TruckService {
	return application.NewTruckService(truckRepository)
}

func TruckRepositoryFactory(db *gorm.DB) application.ITruckRepository {
	return NewTruckRepository(db)
}

func HandlerFactory(ts *application.TruckService) http.Handler {
	return api.NewRouter(api.NewTruckHandler(ts)).GetRoutes()
}

func InitializeWebServer() error {
	config, cErr := NewConfig()
	if cErr != nil {
		return cErr
	}

	db, dErr := DatabaseFactory(config)
	if dErr != nil {
		return dErr
	}

	ts := TruckServiceFactory(TruckRepositoryFactory(db))

	return http.ListenAndServe(config.AppPort, HandlerFactory(ts))
}
