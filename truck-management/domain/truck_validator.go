package domain

func IsValidTruck(newTruck Truck, possibleExistingTruck Truck) error {
	if newTruck.LicensePlate == possibleExistingTruck.LicensePlate {
		return NewConflict("there is already a truck registered with this License Plate")
	}

	if newTruck.EldID == possibleExistingTruck.EldID {
		return NewConflict("there is already a truck registered with this ELD id")
	}

	return nil
}
