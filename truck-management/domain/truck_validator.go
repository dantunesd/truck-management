package domain

type LicensePlateChecker interface {
	IsAlreadyInUse(string) bool
}

type EldChecker interface {
	IsAlreadyInUse(string) bool
}

type TruckValidator struct {
	licensePlateChecker LicensePlateChecker
	eldChecker          EldChecker
}

func NewTruckValidator(licensePlateChecker LicensePlateChecker, eldChecker EldChecker) *TruckValidator {
	return &TruckValidator{
		licensePlateChecker: licensePlateChecker,
		eldChecker:          eldChecker,
	}
}

func (t *TruckValidator) IsValidTruck(newTruck Truck) error {
	if t.licensePlateChecker.IsAlreadyInUse(newTruck.LicensePlate) {
		return NewConflict("there is already a truck registered with this License Plate")
	}

	if t.eldChecker.IsAlreadyInUse(newTruck.EldID) {
		return NewConflict("there is already a truck registered with this ELD id")
	}

	return nil
}
