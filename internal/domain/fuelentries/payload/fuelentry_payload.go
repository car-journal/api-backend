// Package fuelentrypayload handles payload for fuel entry domain
package fuelentrypayload

import (
	"github.com/car-journal/api-backend/lib/nullable"
)

type CreatePayload struct {
	UserID           string  `json:"user_id"`
	CarID            string  `json:"car_id"`
	FuelType         string  `json:"fuel_type"`
	FuelBrand        string  `json:"fuel_brand"`
	FuelName         string  `json:"fuel_name"`
	FuelPrice        float64 `json:"fuel_price"`
	FuelUnit         string  `json:"fuel_unit"`
	DistanceTraveled float64 `json:"distance_traveled"`
	VolumeFilled     float64 `json:"volume_filled"`
	Notes            *string `json:"notes"`
}

type UpdatePayload struct {
	ID               string          `json:"id"`
	UserID           string          `json:"user_id"`
	CarID            string          `json:"car_id"`
	FuelType         *string         `json:"fuel_type"`
	FuelBrand        *string         `json:"fuel_brand"`
	FuelName         *string         `json:"fuel_name"`
	FuelPrice        *float64        `json:"fuel_price"`
	FuelUnit         *string         `json:"fuel_unit"`
	DistanceTraveled *float64        `json:"distance_traveled"`
	VolumeFilled     *float64        `json:"volume_filled"`
	Notes            nullable.String `json:"notes"`
}
