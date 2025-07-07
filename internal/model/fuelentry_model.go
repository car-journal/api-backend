package internalmodel

import "github.com/car-journal/api-backend/lib/uuid"

const FuelEntryTableName = "fuel_entries"

type FuelEntry struct {
	BaseModel

	CarID               uuid.UUID `json:"car_id"`
	FuelType            string    `json:"fuel_type"`
	FuelBrand           string    `json:"fuel_brand"`
	FuelName            string    `json:"fuel_name"`
	FuelPrice           float64   `json:"fuel_price"`
	FuelUnit            string    `json:"fuel_unit"`
	DistanceTraveled    float64   `json:"distance_traveled"`
	VolumeFilled        float64   `json:"volume_filled"`
	TotalPrice          float64   `json:"total_price"`
	FuelConsumptionRate float64   `json:"fuel_consumption"`
	Notes               *string   `json:"notes"`
}
