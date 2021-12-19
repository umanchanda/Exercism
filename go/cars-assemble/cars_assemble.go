package cars

const ProductionRate int = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0
	} else if speed >= 1 && speed <= 4 {
		return 1.0
	} else if speed >= 5 && speed <= 8 {
		return 0.9
	} else {
		return 0.77
	}
	// panic("SuccessRate not implemented")
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * float64(ProductionRate*speed)
	// panic("CalculateProductionRatePerHour not implemented")
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
	// panic("CalculateProductionRatePerMinute not implemented")
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	ProductionPerMinute := CalculateProductionRatePerHour(speed)
	if float64(ProductionPerMinute) > limit {
		return limit
	} else {
		return ProductionPerMinute
	}
	// panic("CalculateLimitedProductionRatePerHour not implemented")
}
