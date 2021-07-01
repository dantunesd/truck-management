package infrastructure

type LicensePlateChecker struct {
}

func (l *LicensePlateChecker) IsAlreadyInUse(licensePlate string) bool {
	// repository usage goes here
	return false
}
