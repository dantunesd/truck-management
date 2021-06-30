package domain

type TruckValidator func(newTruck Truck, existingTruck Truck) error

func IsValidTruck(newTruck Truck, existingTruck Truck) error {
	if newTruck.LicensePlate == existingTruck.LicensePlate {
		return NewConflict("there is already a truck registered with this License Plate")
	}

	if newTruck.EldID == existingTruck.EldID {
		return NewConflict("there is already a truck registered with this ELD id")
	}

	return nil
}
