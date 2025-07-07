// Package fuelentrypayload handles payload for fuel entry domain
package fuelentrypayload

import (
	"github.com/car-journal/api-backend/lib/nullable"
)

type CreatePayload struct {
	OdometerReading  float64 `json:"odometer_reading" validate:"required"`
	ReadingUnit      string  `json:"reading_unit" validate:"required"`
	UserID           string  `json:"user_id"`
	CarID            string  `json:"car_id" validate:"required"`
	FuelType         string  `json:"fuel_type" validate:"required"`
	FuelBrand        string  `json:"fuel_brand" validate:"required"`
	FuelName         string  `json:"fuel_name" validate:"required"`
	FuelPrice        float64 `json:"fuel_price" validate:"required"`
	FuelUnit         string  `json:"fuel_unit" validate:"required"`
	DistanceTraveled float64 `json:"distance_traveled" validate:"required"`
	VolumeFilled     float64 `json:"volume_filled" validate:"required"`
	Notes            *string `json:"notes"`
}

type UpdatePayload struct {
	OdometerReading  *float64        `json:"odometer_reading"`
	ReadingUnit      *string         `json:"reading_unit" validate:"required_with=OdometerReading"`
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
