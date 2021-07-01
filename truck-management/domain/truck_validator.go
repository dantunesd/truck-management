package domain

type LicensePlateChecker interface {
	IsAlreadyInUse(string) bool
}

type EldChecker interface {
	IsAlreadyInUse(string) bool
}

type TruckValidator struct {
	LicensePlateChecker LicensePlateChecker
	EldChecker          EldChecker
}

func (t *TruckValidator) IsValidTruck(newTruck Truck) error {
	if t.LicensePlateChecker.IsAlreadyInUse(newTruck.LicensePlate) {
		return NewConflict("there is already a truck registered with this License Plate")
	}

	if t.EldChecker.IsAlreadyInUse(newTruck.EldID) {
		return NewConflict("there is already a truck registered with this ELD id")
	}

	return nil
}
