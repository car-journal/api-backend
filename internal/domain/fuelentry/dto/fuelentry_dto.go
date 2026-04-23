// Package fuelentrydto handles const for odometer entry domain
package fuelentrydto

import (
	internalmodel "github.com/car-journal/api-backend/internal/model"
	"github.com/car-journal/api-backend/lib/uuid"
)

type FuelEntry struct {
	ID                  uuid.UUID `json:"id"`
	FuelBrand           string    `json:"fuel_type"`
	FuelName            string    `json:"fuel_name"`
	FuelUnit            string    `json:"fuel_unit"`
	VolumeFilled        float64   `json:"volume_filled"`
	TotalPrice          float64   `json:"total_price"`
	FuelConsumptionRate float64   `json:"fuel_consumption_rate"`
}

type FuelEntryWithOdomoeterReading struct {
	*internalmodel.FuelEntry
	OdometerReading float64 `json:"odometer_reading"`
}

type ConsumptionRateStats struct {
	TotalFuelCost   float64 `gorm:"column:total_fuel_cost"       json:"total_fuel_cost"`
	AllTimeAvgRate  float64 `gorm:"column:all_time_avg_rate"     json:"all_time_avg_rate"`
	CurrentRate     float64 `gorm:"column:current_rate"          json:"current_rate"`
	PreviousAvgRate float64 `gorm:"column:previous_avg_rate"     json:"previous_avg_rate"`
}
