package infrastructure

import (
	"fmt"
	"net/http"
	"truck-management/truck-management/api"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseFactory(config *Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.DBConn), &gorm.Config{})
}

func TruckServiceFactory(truckRepository application.ITruckRepository) *application.TruckService {
	return application.NewTruckService(
		truckRepository,
		TruckValidatorFactory(),
	)
}

func TruckRepositoryFactory(db *gorm.DB) application.ITruckRepository {
	return NewTruckRepository(db)
}

func TruckValidatorFactory() application.ITruckValidator {
	return domain.NewTruckValidator(
		LicensePlateCheckerFactory(),
		EldCheckerFactory(),
	)
}

func LicensePlateCheckerFactory() domain.LicensePlateChecker {
	return &LicensePlateChecker{}
}

func EldCheckerFactory() domain.EldChecker {
	return &EldChecker{}
}

func HandlerFactory() http.Handler {
	config, cerr := NewConfig()
	db, derr := DatabaseFactory(config)
	truckRepository := TruckRepositoryFactory(db)

	fmt.Println(cerr, derr)

	return api.NewRouter(api.NewTruckHandler(TruckServiceFactory(truckRepository))).GetRoutes()
}
