package internalmodel

import (
	"time"

	"github.com/car-journal/api-backend/lib/uuid"
)

const MaintenanceEntryTableName = "maintenance_entries"

type MaintenanceEntry struct {
	BaseModel

	CarID           uuid.UUID `json:"car_id"`
	OdometerEntryID uuid.UUID `json:"odometer_entry_id"`
	CategoryID      uuid.UUID `json:"category_id"`
	Brand           string    `json:"brand"`
	Name            string    `json:"name"`
	Price           float64   `json:"price"`
	PerformedAt     time.Time `json:"performed_at"`
	Notes           *string   `json:"notes"`
}
