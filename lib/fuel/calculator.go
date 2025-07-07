// Package fuel handles all necessary calculation or data modification for fuel domain
package fuel

func CalculateFuelConsumptionRate(distanceTraveled float64, volumeFilled float64) float64 {
	return distanceTraveled / volumeFilled
}

func CalculateTotalPrice(pricePerUnit float64, volumeFilled float64) float64 {
	return pricePerUnit * volumeFilled
}
