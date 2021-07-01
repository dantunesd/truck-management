package infrastructure

import (
	"truck-management/truck-management/application"
	"truck-management/truck-management/domain"
)

func CreateTruckServiceFactory() *application.CreateTruckService {
	return &application.CreateTruckService{
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
