// Package maintenanceentrydto contains data transfer objects related to maintenance entries.
package maintenanceentrydto

import internalmodel "github.com/car-journal/api-backend/internal/model"

type MaintenanceEntryWithOdometerReading struct {
	*internalmodel.MaintenanceEntry
	OdometerReading float64 `json:"odometer_reading"`
}
