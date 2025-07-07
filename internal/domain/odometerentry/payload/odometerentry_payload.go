// Package odometerentrypayload handles payload for odometer entry domain
package odometerentrypayload

import (
	"github.com/car-journal/api-backend/lib/uuid"
)

type CreatePayload struct {
	UserID          string    `json:"user_id"`
	CarID           uuid.UUID `json:"car_id"`
	OdometerReading float64   `json:"odometer_reading" validate:"required"`
	ReadingUnit     string    `json:"reading_unit" validate:"required"`
}

type UpdatePayload struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	CarID           uuid.UUID `json:"car_id"`
	OdometerReading *float64  `json:"odometer_reading"`
	ReadingUnit     *string   `json:"reading_unit" validate:"required_with=OdometerReading"`
}
