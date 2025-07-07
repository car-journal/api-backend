// Package fuelentrydto handles const for odometer entry domain
package fuelentrydto

import "github.com/car-journal/api-backend/lib/uuid"

type FuelEntry struct {
	ID                  uuid.UUID `json:"id"`
	FuelBrand           string    `json:"fuel_type"`
	FuelName            string    `json:"fuel_name"`
	FuelUnit            string    `json:"fuel_unit"`
	VolumeFilled        float64   `json:"volume_filled"`
	TotalPrice          float64   `json:"total_price"`
	FuelConsumptionRate float64   `json:"fuel_consumption_rate"`
}
