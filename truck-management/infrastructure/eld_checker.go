package infrastructure

type EldChecker struct {
}

func (l *EldChecker) IsAlreadyInUse(eldId string) bool {
	// repository usage goes here
	return false
}
