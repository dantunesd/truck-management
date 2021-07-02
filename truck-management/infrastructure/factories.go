package infrastructure

import (
	"net/http"
	"truck-management/truck-management/api"
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"
)

func TruckServiceFactory() *application.TruckService {
	return application.NewTruckService(
		TruckRepositoryFactory(),
		TruckValidatorFactory(),
	)
}

func TruckRepositoryFactory() application.ITruckRepository {
	return NewTruckRepository()
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
	return api.NewRouter(api.NewTruckHandler(TruckServiceFactory())).GetRoutes()
}
