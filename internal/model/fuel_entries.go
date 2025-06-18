package internalmodel

import "github.com/car-journal/lib/uuid"

const FuelEntriesTableName = "fuel_entries"

type FuelEntries struct {
	BaseModel

	CarID            uuid.UUID `json:"car_id"`
	FuelType         string    `json:"fuel_type"`
	FuelBrand        string    `json:"fuel_brand"`
	FuelName         string    `json:"fuel_name"`
	FuelPrice        float64   `json:"fuel_price"`
	FuelUnit         string    `json:"fuel_unit"`
	DistanceTraveled float64   `json:"distance_traveled"`
	VolumeFilled     float64   `json:"volume_filled"`
	TotalPrice       float64   `json:"total_price"`
	FuelConsumption  float64   `json:"fuel_consumption"`
	Notes            *string   `json:"notes"`
}
