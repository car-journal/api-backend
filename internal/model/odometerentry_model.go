package internalmodel

import "github.com/car-journal/api-backend/lib/uuid"

const OdometerEntryTableName = "odometer_entries"

type OdometerEntry struct {
	BaseModel

	CarID           uuid.UUID `json:"car_id"`
	OdometerReading float64   `json:"odometer_reading"`
	ReadingUnit     string    `json:"reading_unit"`
}
