package infrastructure

import (
	"net/http"
	"truck-management/truck-management/api"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"
)

func TruckServiceFactory() *application.TruckService {
	return &application.TruckService{
		TruckRepository: TruckRepositoryFactory(),
		TruckValidator:  TruckValidatorFactory(),
	}
}

func TruckRepositoryFactory() application.ITruckRepository {
	return &TruckRepository{}
}

func TruckValidatorFactory() application.ITruckValidator {
	return &domain.TruckValidator{
		LicensePlateChecker: LicensePlateCheckerFactory(),
		EldChecker:          EldCheckerFactory(),
	}
}

func LicensePlateCheckerFactory() domain.LicensePlateChecker {
	return &LicensePlateChecker{}
}

func EldCheckerFactory() domain.EldChecker {
	return &EldChecker{}
}

func HandlerFactory() http.Handler {
	return api.Router(TruckServiceFactory())
}
