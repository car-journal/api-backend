// Package maintenanceentrypayload handles payload for maintenance entry domain
package maintenanceentrypayload

import (
	"time"

	"github.com/car-journal/api-backend/lib/filter"
	"github.com/car-journal/api-backend/lib/nullable"
)

type CreatePayload struct {
	CarID           string    `json:"car_id"`
	CategoryID      string    `json:"category_id" validate:"required"`
	OdometerReading float64   `json:"odometer_reading" validate:"required"`
	ReadingUnit     string    `json:"reading_unit" validate:"required"`
	Brand           string    `json:"brand" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Price           float64   `json:"price" validate:"required"`
	PerformedAt     time.Time `json:"performed_at" validate:"required"`
	Notes           *string   `json:"notes"`
}

type ListPayload struct {
	CarID       string       `json:"car_id"`
	Name        string       `json:"name"`
	CategoryIDs []string     `json:"category_ids"`
	Sorts       []string     `json:"sorts"`
	PageParams  *filter.Page `json:"page_params"`
}

type UpdatePayload struct {
	ID              string          `json:"id"`
	CategoryID      *string         `json:"category_id"`
	OdometerReading *float64        `json:"odometer_reading"`
	ReadingUnit     *string         `json:"reading_unit" validate:"required_with=OdometerReading"`
	Brand           *string         `json:"brand"`
	Name            *string         `json:"name"`
	Price           *float64        `json:"price"`
	PerformedAt     nullable.Time   `json:"performed_at"`
	Notes           nullable.String `json:"notes"`
}
